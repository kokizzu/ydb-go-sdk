package table

import (
	"bytes"
	"context"
	"io"
	"sync"
	"time"

	"github.com/yandex-cloud/ydb-go-sdk"
	"github.com/yandex-cloud/ydb-go-sdk/api/grpc/Ydb_Table_V1"
	"github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb"
	"github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb_Table"
	"github.com/yandex-cloud/ydb-go-sdk/internal"
	"github.com/yandex-cloud/ydb-go-sdk/internal/cache/lru"
)

var DefaultMaxQueryCacheSize = 1000

// Client contains logic of creation of ydb table sessions.
type Client struct {
	Driver ydb.Driver
	Trace  ClientTrace

	// MaxQueryCacheSize limits maximum number of queries which able to live in
	// cache. Note that cache is not shared across sessions.
	//
	// If MaxQueryCacheSize equal to zero, then the DefaultMaxQueryCacheSize is used.
	//
	// If MaxQueryCacheSize is less than zero, then client not used query cache.
	MaxQueryCacheSize int
}

// CreateSession creates new session instance.
// Unused sessions must be destroyed.
func (t *Client) CreateSession(ctx context.Context) (s *Session, err error) {
	t.traceCreateSessionStart(ctx)
	start := time.Now()
	defer func() {
		t.traceCreateSessionDone(ctx, s, time.Since(start), err)
	}()
	var (
		req Ydb_Table.CreateSessionRequest
		res Ydb_Table.CreateSessionResult
	)
	var endpointInfo ydb.EndpointInfo
	endpointInfo, err = t.Driver.Call(
		ctx,
		internal.Wrap(
			Ydb_Table_V1.CreateSession,
			&req,
			&res,
		),
	)
	if err != nil {
		return nil, err
	}
	s = &Session{
		ID:           res.SessionId,
		endpointInfo: endpointInfo,
		c:            *t,
		qcache:       lru.New(t.cacheSize()),
	}
	return
}

func (t *Client) cacheSize() int {
	if t.MaxQueryCacheSize == 0 {
		return DefaultMaxQueryCacheSize
	}
	return t.MaxQueryCacheSize
}

// Session represents a single table API session.
//
// Session methods are not goroutine safe. Simultaneous execution of requests
// are forbidden within a single session.
//
// Note that after Session is no longer needed it should be destroyed by
// Close() call.
type Session struct {
	ID           string
	endpointInfo ydb.EndpointInfo

	c Client

	qcache lru.Cache
	qhash  queryHasher

	closeMux sync.Mutex
	closed   bool
	onClose  []func()
}

func (s *Session) OnClose(cb func()) {
	s.closeMux.Lock()
	defer s.closeMux.Unlock()
	if s.closed {
		return
	}
	s.onClose = append(s.onClose, cb)
}

func (s *Session) Close(ctx context.Context) (err error) {
	s.closeMux.Lock()
	defer s.closeMux.Unlock()
	if s.closed {
		return nil
	}
	s.closed = true
	s.c.traceDeleteSessionStart(ctx, s)
	start := time.Now()
	defer func() {
		for _, cb := range s.onClose {
			cb()
		}
		s.c.traceDeleteSessionDone(ctx, s, time.Since(start), err)
	}()
	req := Ydb_Table.DeleteSessionRequest{
		SessionId: s.ID,
	}
	_, err = s.c.Driver.Call(
		ydb.WithEndpointInfo(
			ctx,
			s.endpointInfo,
		),
		internal.Wrap(
			Ydb_Table_V1.DeleteSession,
			&req,
			nil,
		),
	)
	return err
}

func (s *Session) Address() string {
	if s.endpointInfo != nil {
		return s.endpointInfo.Address()
	}
	return ""
}

// KeepAlive keeps idle session alive.
func (s *Session) KeepAlive(ctx context.Context) (info SessionInfo, err error) {
	s.c.traceKeepAliveStart(ctx, s)
	defer func() {
		s.c.traceKeepAliveDone(ctx, s, info, err)
	}()
	var res Ydb_Table.KeepAliveResult
	req := Ydb_Table.KeepAliveRequest{
		SessionId: s.ID,
	}
	_, err = s.c.Driver.Call(
		ydb.WithEndpointInfoAndPolicy(
			ctx,
			s.endpointInfo,
			ydb.ConnUseEndpoint,
		),
		internal.Wrap(
			Ydb_Table_V1.KeepAlive,
			&req,
			&res,
		),
	)
	if err != nil {
		return
	}
	switch res.SessionStatus {
	case Ydb_Table.KeepAliveResult_SESSION_STATUS_READY:
		info.Status = SessionReady
	case Ydb_Table.KeepAliveResult_SESSION_STATUS_BUSY:
		info.Status = SessionBusy
	}
	return
}

// CreateTable creates table at given path with given options.
func (s *Session) CreateTable(ctx context.Context, path string, opts ...CreateTableOption) (err error) {
	req := Ydb_Table.CreateTableRequest{
		SessionId: s.ID,
		Path:      path,
	}
	for _, opt := range opts {
		opt((*createTableDesc)(&req))
	}
	_, err = s.c.Driver.Call(
		ydb.WithEndpointInfo(
			ctx,
			s.endpointInfo,
		),
		internal.Wrap(
			Ydb_Table_V1.CreateTable,
			&req,
			nil,
		),
	)
	return err
}

// DescribeTable describes table at given path.
func (s *Session) DescribeTable(ctx context.Context, path string, opts ...DescribeTableOption) (desc Description, err error) {
	var res Ydb_Table.DescribeTableResult
	req := Ydb_Table.DescribeTableRequest{
		SessionId: s.ID,
		Path:      path,
	}
	for _, opt := range opts {
		opt((*describeTableDesc)(&req))
	}
	_, err = s.c.Driver.Call(
		ydb.WithEndpointInfo(
			ctx,
			s.endpointInfo,
		),
		internal.Wrap(
			Ydb_Table_V1.DescribeTable,
			&req,
			&res,
		),
	)
	if err != nil {
		return desc, err
	}

	cs := make([]Column, len(res.Columns))
	for i, c := range res.Columns {
		cs[i] = Column{
			Name:   c.Name,
			Type:   internal.TypeFromYDB(c.Type),
			Family: c.Family,
		}
	}

	rs := make([]KeyRange, len(res.ShardKeyBounds)+1)
	var last ydb.Value
	for i, b := range res.ShardKeyBounds {
		if last != nil {
			rs[i].From = last
		}

		bound := internal.ValueFromYDB(b.Type, b.Value)
		rs[i].To = bound

		last = bound
	}
	if last != nil {
		i := len(rs) - 1
		rs[i].From = last
	}

	var stats *TableStats
	if res.TableStats != nil {
		resStats := res.TableStats
		partStats := make([]PartitionStats, len(res.TableStats.PartitionStats))
		for i, v := range res.TableStats.PartitionStats {
			partStats[i].RowsEstimate = v.RowsEstimate
			partStats[i].StoreSize = v.StoreSize
		}
		var creationTime, modificationTime time.Time
		if resStats.CreationTime.GetSeconds() != 0 {
			creationTime = time.Unix(resStats.CreationTime.GetSeconds(), int64(resStats.CreationTime.GetNanos()))
		}
		if resStats.ModificationTime.GetSeconds() != 0 {
			modificationTime = time.Unix(resStats.ModificationTime.GetSeconds(), int64(resStats.ModificationTime.GetNanos()))
		}

		stats = &TableStats{
			PartitionStats:   partStats,
			RowsEstimate:     resStats.RowsEstimate,
			StoreSize:        resStats.StoreSize,
			Partitions:       resStats.Partitions,
			CreationTime:     creationTime,
			ModificationTime: modificationTime,
		}
	}

	cf := make([]ColumnFamily, len(res.ColumnFamilies))
	for i, c := range res.ColumnFamilies {
		cf[i] = columnFamily(c)
	}

	attrs := make(map[string]string, len(res.Attributes))
	for k, v := range res.Attributes {
		attrs[k] = v
	}

	return Description{
		Name:                 res.Self.Name,
		PrimaryKey:           res.PrimaryKey,
		Columns:              cs,
		KeyRanges:            rs,
		Stats:                stats,
		ColumnFamilies:       cf,
		Attributes:           attrs,
		ReadReplicaSettings:  readReplicasSettings(res.GetReadReplicasSettings()),
		StorageSettings:      storageSettings(res.GetStorageSettings()),
		KeyBloomFilter:       internal.FeatureFlagFromYDB(res.GetKeyBloomFilter()),
		PartitioningSettings: partitioningSettings(res.GetPartitioningSettings()),
		TTLSettings:          ttlSettings(res.GetTtlSettings()),
		TimeToLiveSettings:   timeToLiveSettings(res.GetTtlSettings()),
	}, nil
}

// DropTable drops table at given path with given options.
func (s *Session) DropTable(ctx context.Context, path string, opts ...DropTableOption) (err error) {
	req := Ydb_Table.DropTableRequest{
		SessionId: s.ID,
		Path:      path,
	}
	for _, opt := range opts {
		opt((*dropTableDesc)(&req))
	}
	_, err = s.c.Driver.Call(
		ydb.WithEndpointInfo(
			ctx,
			s.endpointInfo,
		),
		internal.Wrap(
			Ydb_Table_V1.DropTable,
			&req,
			nil,
		),
	)
	return err
}

// AlterTable modifies schema of table at given path with given options.
func (s *Session) AlterTable(ctx context.Context, path string, opts ...AlterTableOption) (err error) {
	req := Ydb_Table.AlterTableRequest{
		SessionId: s.ID,
		Path:      path,
	}
	for _, opt := range opts {
		opt((*alterTableDesc)(&req))
	}
	_, err = s.c.Driver.Call(
		ydb.WithEndpointInfo(
			ctx,
			s.endpointInfo,
		),
		internal.Wrap(
			Ydb_Table_V1.AlterTable,
			&req,
			nil,
		),
	)
	return err
}

// CopyTable creates copy of table at given path.
func (s *Session) CopyTable(ctx context.Context, dst, src string, _ ...CopyTableOption) (err error) {
	req := Ydb_Table.CopyTableRequest{
		SessionId:       s.ID,
		SourcePath:      src,
		DestinationPath: dst,
	}
	_, err = s.c.Driver.Call(
		ydb.WithEndpointInfo(
			ctx,
			s.endpointInfo,
		),
		internal.Wrap(
			Ydb_Table_V1.CopyTable,
			&req,
			nil,
		),
	)
	return err
}

// DataQueryExplanation is a result of ExplainDataQuery call.
type DataQueryExplanation struct {
	AST  string
	Plan string
}

// Explain explains data query represented by text.
func (s *Session) Explain(ctx context.Context, query string) (exp DataQueryExplanation, err error) {
	var res Ydb_Table.ExplainQueryResult
	req := Ydb_Table.ExplainDataQueryRequest{
		SessionId: s.ID,
		YqlText:   query,
	}
	_, err = s.c.Driver.Call(
		ydb.WithEndpointInfo(
			ctx,
			s.endpointInfo,
		),
		internal.Wrap(
			Ydb_Table_V1.ExplainDataQuery,
			&req,
			&res,
		),
	)
	if err != nil {
		return
	}
	return DataQueryExplanation{
		AST:  res.QueryAst,
		Plan: res.QueryPlan,
	}, nil
}

// Statement is a prepared statement. Like a single Session, it is not safe for
// concurrent use by multiple goroutines.
type Statement struct {
	session *Session
	query   *DataQuery
	qhash   queryHash
	params  map[string]*Ydb.Type
}

// Execute executes prepared data query.
func (s *Statement) Execute(
	ctx context.Context, tx *TransactionControl,
	params *QueryParameters,
	opts ...ExecuteDataQueryOption,
) (
	txr *Transaction, r *Result, err error,
) {
	s.session.c.traceExecuteDataQueryStart(ctx, s.session, tx, s.query, params)
	defer func() {
		s.session.c.traceExecuteDataQueryDone(ctx, s.session, tx, s.query, params, true, txr, r, err)
	}()
	return s.execute(ctx, tx, params, opts...)
}

// execute executes prepared query without any tracing.
func (s *Statement) execute(
	ctx context.Context, tx *TransactionControl,
	params *QueryParameters,
	opts ...ExecuteDataQueryOption,
) (
	txr *Transaction, r *Result, err error,
) {
	_, res, err := s.session.executeDataQuery(ctx, tx, s.query, params, opts...)
	if ydb.IsOpError(err, ydb.StatusNotFound) {
		s.session.qcache.Remove(s.qhash)
	}
	if err != nil {
		return nil, nil, err
	}
	return s.session.executeQueryResult(res)
}

func (s *Statement) NumInput() int {
	return len(s.params)
}

// Prepare prepares data query within session s.
func (s *Session) Prepare(
	ctx context.Context, query string,
) (
	stmt *Statement, err error,
) {
	var (
		cached bool
		q      *DataQuery
	)
	s.c.tracePrepareDataQueryStart(ctx, s, query)
	defer func() {
		s.c.tracePrepareDataQueryDone(ctx, s, query, q, cached, err)
	}()

	cacheKey := s.qhash.hash(query)
	stmt, cached = s.getQueryFromCache(cacheKey)
	if cached {
		q = stmt.query
		return stmt, nil
	}

	var res Ydb_Table.PrepareQueryResult
	req := Ydb_Table.PrepareDataQueryRequest{
		SessionId: s.ID,
		YqlText:   query,
	}
	_, err = s.c.Driver.Call(
		ydb.WithEndpointInfo(
			ctx,
			s.endpointInfo,
		),
		internal.Wrap(
			Ydb_Table_V1.PrepareDataQuery,
			&req,
			&res,
		),
	)
	if err != nil {
		return nil, err
	}

	q = new(DataQuery)
	q.initPrepared(res.QueryId)
	stmt = &Statement{
		session: s,
		query:   q,
		qhash:   cacheKey,
		params:  res.ParametersTypes,
	}
	s.addQueryToCache(cacheKey, stmt)

	return stmt, nil
}

func (s *Session) getQueryFromCache(key queryHash) (*Statement, bool) {
	v, cached := s.qcache.Get(key)
	if cached {
		return v.(*Statement), true
	}
	return nil, false
}

func (s *Session) addQueryToCache(key queryHash, stmt *Statement) {
	s.qcache.Add(key, stmt)
}

// Execute executes given data query represented by text.
func (s *Session) Execute(
	ctx context.Context,
	tx *TransactionControl,
	query string,
	params *QueryParameters,
	opts ...ExecuteDataQueryOption,
) (
	txr *Transaction, r *Result, err error,
) {
	q := new(DataQuery)
	q.initFromText(query)

	var cached bool
	s.c.traceExecuteDataQueryStart(ctx, s, tx, q, params)
	defer func() {
		s.c.traceExecuteDataQueryDone(ctx, s, tx, q, params, cached, txr, r, err)
	}()

	cacheKey := s.qhash.hash(query)
	stmt, cached := s.getQueryFromCache(cacheKey)
	if cached {
		// Supplement q with ID for tracing.
		q.initPreparedText(query, stmt.query.ID())
		return stmt.execute(ctx, tx, params, opts...)
	}
	req, res, err := s.executeDataQuery(ctx, tx, q, params, opts...)
	if err != nil {
		return nil, nil, err
	}
	if keepInCache(req) && res.QueryMeta != nil {
		queryID := res.QueryMeta.Id
		// Supplement q with ID for tracing.
		q.initPreparedText(query, queryID)
		// Create new DataQuery instead of q above to not store the whole query
		// string within statement.
		subq := new(DataQuery)
		subq.initPrepared(queryID)
		stmt = &Statement{
			session: s,
			query:   subq,
			qhash:   cacheKey,
			params:  res.QueryMeta.ParametersTypes,
		}
		s.addQueryToCache(cacheKey, stmt)
	}

	return s.executeQueryResult(res)
}

func keepInCache(req *Ydb_Table.ExecuteDataQueryRequest) bool {
	p := req.QueryCachePolicy
	return p != nil && p.KeepInCache
}

// executeQueryResult returns Transaction and Result built from received
// result.
func (s *Session) executeQueryResult(res *Ydb_Table.ExecuteQueryResult) (*Transaction, *Result, error) {
	t := &Transaction{
		id: res.TxMeta.Id,
		s:  s,
	}
	r := &Result{
		sets:  res.ResultSets,
		stats: res.QueryStats,
	}
	return t, r, nil
}

// executeDataQuery executes data query.
func (s *Session) executeDataQuery(
	ctx context.Context, tx *TransactionControl,
	query *DataQuery, params *QueryParameters,
	opts ...ExecuteDataQueryOption,
) (
	req *Ydb_Table.ExecuteDataQueryRequest,
	res *Ydb_Table.ExecuteQueryResult,
	err error,
) {
	res = new(Ydb_Table.ExecuteQueryResult)
	req = &Ydb_Table.ExecuteDataQueryRequest{
		SessionId:  s.ID,
		TxControl:  &tx.desc,
		Parameters: params.params(),
		Query:      &query.query,
	}
	for _, opt := range opts {
		opt((*executeDataQueryDesc)(req))
	}
	if m, _ := ydb.ContextOperationMode(ctx); m == ydb.OperationModeUnknown {
		ctx = ydb.WithOperationMode(ctx, ydb.OperationModeSync)
	}
	_, err = s.c.Driver.Call(
		ydb.WithEndpointInfo(
			ctx,
			s.endpointInfo,
		),
		internal.Wrap(
			Ydb_Table_V1.ExecuteDataQuery,
			req,
			res,
		),
	)
	return
}

// ExecuteSchemeQuery executes scheme query.
func (s *Session) ExecuteSchemeQuery(
	ctx context.Context, query string,
	opts ...ExecuteSchemeQueryOption,
) (err error) {
	req := Ydb_Table.ExecuteSchemeQueryRequest{
		SessionId: s.ID,
		YqlText:   query,
	}
	for _, opt := range opts {
		opt((*executeSchemeQueryDesc)(&req))
	}
	_, err = s.c.Driver.Call(
		ydb.WithEndpointInfo(
			ctx,
			s.endpointInfo,
		),
		internal.Wrap(
			Ydb_Table_V1.ExecuteSchemeQuery,
			&req,
			nil,
		),
	)
	return err
}

// DescribeTableOptions describes supported table options.
func (s *Session) DescribeTableOptions(ctx context.Context) (desc TableOptionsDescription, err error) {
	var res Ydb_Table.DescribeTableOptionsResult
	req := Ydb_Table.DescribeTableOptionsRequest{}
	_, err = s.c.Driver.Call(
		ydb.WithEndpointInfo(
			ctx,
			s.endpointInfo,
		),
		internal.Wrap(
			Ydb_Table_V1.DescribeTableOptions,
			&req,
			&res,
		),
	)
	if err != nil {
		return
	}
	{
		xs := make([]TableProfileDescription, len(res.TableProfilePresets))
		for i, p := range res.TableProfilePresets {
			xs[i] = TableProfileDescription{
				Name:   p.Name,
				Labels: p.Labels,

				DefaultStoragePolicy:      p.DefaultStoragePolicy,
				DefaultCompactionPolicy:   p.DefaultCompactionPolicy,
				DefaultPartitioningPolicy: p.DefaultPartitioningPolicy,
				DefaultExecutionPolicy:    p.DefaultExecutionPolicy,
				DefaultReplicationPolicy:  p.DefaultReplicationPolicy,
				DefaultCachingPolicy:      p.DefaultCachingPolicy,

				AllowedStoragePolicies:      p.AllowedStoragePolicies,
				AllowedCompactionPolicies:   p.AllowedCompactionPolicies,
				AllowedPartitioningPolicies: p.AllowedPartitioningPolicies,
				AllowedExecutionPolicies:    p.AllowedExecutionPolicies,
				AllowedReplicationPolicies:  p.AllowedReplicationPolicies,
				AllowedCachingPolicies:      p.AllowedCachingPolicies,
			}
		}
		desc.TableProfilePresets = xs
	}
	{
		xs := make([]StoragePolicyDescription, len(res.StoragePolicyPresets))
		for i, p := range res.StoragePolicyPresets {
			xs[i] = StoragePolicyDescription{
				Name:   p.Name,
				Labels: p.Labels,
			}
		}
		desc.StoragePolicyPresets = xs
	}
	{
		xs := make([]CompactionPolicyDescription, len(res.CompactionPolicyPresets))
		for i, p := range res.CompactionPolicyPresets {
			xs[i] = CompactionPolicyDescription{
				Name:   p.Name,
				Labels: p.Labels,
			}
		}
		desc.CompactionPolicyPresets = xs
	}
	{
		xs := make([]PartitioningPolicyDescription, len(res.PartitioningPolicyPresets))
		for i, p := range res.PartitioningPolicyPresets {
			xs[i] = PartitioningPolicyDescription{
				Name:   p.Name,
				Labels: p.Labels,
			}
		}
		desc.PartitioningPolicyPresets = xs
	}
	{
		xs := make([]ExecutionPolicyDescription, len(res.ExecutionPolicyPresets))
		for i, p := range res.ExecutionPolicyPresets {
			xs[i] = ExecutionPolicyDescription{
				Name:   p.Name,
				Labels: p.Labels,
			}
		}
		desc.ExecutionPolicyPresets = xs
	}
	{
		xs := make([]ReplicationPolicyDescription, len(res.ReplicationPolicyPresets))
		for i, p := range res.ReplicationPolicyPresets {
			xs[i] = ReplicationPolicyDescription{
				Name:   p.Name,
				Labels: p.Labels,
			}
		}
		desc.ReplicationPolicyPresets = xs
	}
	{
		xs := make([]CachingPolicyDescription, len(res.CachingPolicyPresets))
		for i, p := range res.CachingPolicyPresets {
			xs[i] = CachingPolicyDescription{
				Name:   p.Name,
				Labels: p.Labels,
			}
		}
		desc.CachingPolicyPresets = xs
	}
	return desc, nil
}

// StreamReadTable reads table at given path with given options.
//
// Note that given ctx controls the lifetime of the whole read, not only this
// StreamReadTable() call; that is, the time until returned result is closed
// via Close() call or fully drained by sequential NextStreamSet() calls.
func (s *Session) StreamReadTable(ctx context.Context, path string, opts ...ReadTableOption) (r *Result, err error) {
	s.c.traceStreamReadTableStart(ctx, s)
	defer func() {
		s.c.traceStreamReadTableDone(ctx, s, r, err)
	}()

	var resp Ydb_Table.ReadTableResponse
	req := Ydb_Table.ReadTableRequest{
		SessionId: s.ID,
		Path:      path,
	}
	for _, opt := range opts {
		opt((*readTableDesc)(&req))
	}

	ctx, cancel := context.WithCancel(ctx)
	var (
		ch   = make(chan *Ydb.ResultSet, 1)
		ce   = new(error)
		once = sync.Once{}
	)
	_, err = s.c.Driver.StreamRead(
		ydb.WithEndpointInfo(
			ctx,
			s.endpointInfo,
		),
		internal.WrapStreamOperation(
			Ydb_Table_V1.StreamReadTable,
			&req,
			&resp,
			func(err error) {
				if err != io.EOF {
					*ce = err
				}
				if err != nil {
					once.Do(func() { close(ch) })
					return
				}
				select {
				case <-ctx.Done():
					once.Do(func() { close(ch) })
				default:
					if result := resp.Result; result != nil {
						if result.ResultSet != nil {
							ch <- resp.Result.ResultSet
						}
					}
				}
			},
		),
	)
	if err != nil {
		cancel()
		return
	}
	r = &Result{
		setCh:       ch,
		setChErr:    ce,
		setChCancel: cancel,
	}
	return r, nil
}

// StreamExecuteScanQuery scan-reads table at given path with given options.
//
// Note that given ctx controls the lifetime of the whole read, not only this
// StreamExecuteScanQuery() call; that is, the time until returned result is closed
// via Close() call or fully drained by sequential NextStreamSet() calls.
func (s *Session) StreamExecuteScanQuery(
	ctx context.Context,
	query string,
	params *QueryParameters,
	opts ...ExecuteScanQueryOption,
) (
	r *Result, err error,
) {
	q := new(DataQuery)
	q.initFromText(query)

	s.c.traceStreamExecuteScanQueryStart(ctx, s, q, params)
	defer func() {
		s.c.traceStreamExecuteScanQueryDone(ctx, s, q, params, r, err)
	}()

	var resp Ydb_Table.ExecuteScanQueryPartialResponse
	req := Ydb_Table.ExecuteScanQueryRequest{
		Query:      &q.query,
		Parameters: params.params(),
		Mode:       Ydb_Table.ExecuteScanQueryRequest_MODE_EXEC, // set default
	}
	for _, opt := range opts {
		opt((*executeScanQueryDesc)(&req))
	}

	ctx, cancel := context.WithCancel(ctx)
	var (
		ch   = make(chan *Ydb.ResultSet, 1)
		ce   = new(error)
		once = sync.Once{}
	)
	_, err = s.c.Driver.StreamRead(
		ydb.WithEndpointInfo(
			ctx,
			s.endpointInfo,
		),
		internal.WrapStreamOperation(
			Ydb_Table_V1.StreamExecuteScanQuery,
			&req,
			&resp,
			func(err error) {
				if err != io.EOF {
					*ce = err
				}
				if err != nil {
					once.Do(func() { close(ch) })
					return
				}
				select {
				case <-ctx.Done():
					once.Do(func() { close(ch) })
				default:
					if result := resp.Result; result != nil {
						if result.ResultSet != nil {
							ch <- resp.Result.ResultSet
						}
						// TODO: something
						// if result.QueryStats != nil {
						// }
					}
				}
			},
		),
	)
	if err != nil {
		cancel()
		return
	}
	r = &Result{
		setCh:       ch,
		setChErr:    ce,
		setChCancel: cancel,
	}
	return r, nil
}

// BulkUpsert uploads given list of ydb struct values to the table.
func (s *Session) BulkUpsert(ctx context.Context, table string, rows ydb.Value) (err error) {
	req := Ydb_Table.BulkUpsertRequest{
		Table: table,
		Rows:  internal.ValueToYDB(rows),
	}
	_, err = s.c.Driver.Call(
		ydb.WithEndpointInfo(
			ctx,
			s.endpointInfo,
		),
		internal.Wrap(
			Ydb_Table_V1.BulkUpsert,
			&req, nil,
		),
	)
	return err
}

// BeginTransaction begins new transaction within given session with given
// settings.
func (s *Session) BeginTransaction(ctx context.Context, tx *TransactionSettings) (x *Transaction, err error) {
	s.c.traceBeginTransactionStart(ctx, s)
	defer func() {
		s.c.traceBeginTransactionDone(ctx, s, x, err)
	}()
	var res Ydb_Table.BeginTransactionResult
	req := Ydb_Table.BeginTransactionRequest{
		SessionId:  s.ID,
		TxSettings: &tx.settings,
	}
	_, err = s.c.Driver.Call(
		ydb.WithEndpointInfo(
			ctx,
			s.endpointInfo,
		),
		internal.Wrap(
			Ydb_Table_V1.BeginTransaction,
			&req,
			&res,
		),
	)
	if err != nil {
		return
	}
	return &Transaction{
		id: res.TxMeta.Id,
		s:  s,
	}, nil
}

// Transaction is a database transaction.
// Hence session methods are not goroutine safe, Transaction is not goroutine
// safe either.
type Transaction struct {
	id string
	s  *Session
	c  *TransactionControl
}

// Execute executes query represented by text within transaction tx.
func (tx *Transaction) Execute(
	ctx context.Context,
	query string, params *QueryParameters,
	opts ...ExecuteDataQueryOption,
) (r *Result, err error) {
	_, r, err = tx.s.Execute(ctx, tx.txc(), query, params, opts...)
	return
}

// ExecuteStatement executes prepared statement stmt within transaction tx.
func (tx *Transaction) ExecuteStatement(
	ctx context.Context,
	stmt *Statement, params *QueryParameters,
	opts ...ExecuteDataQueryOption,
) (r *Result, err error) {
	_, r, err = stmt.Execute(ctx, tx.txc(), params, opts...)
	return
}

// Deprecated: Use CommitTx instead
// Commit commits specified active transaction.
func (tx *Transaction) Commit(ctx context.Context) (err error) {
	tx.s.c.traceCommitTransactionStart(ctx, tx)
	defer func() {
		tx.s.c.traceCommitTransactionDone(ctx, tx, err)
	}()
	req := Ydb_Table.CommitTransactionRequest{
		SessionId: tx.s.ID,
		TxId:      tx.id,
	}
	_, err = tx.s.c.Driver.Call(
		ydb.WithEndpointInfo(
			ctx,
			tx.s.endpointInfo,
		),
		internal.Wrap(
			Ydb_Table_V1.CommitTransaction,
			&req,
			nil,
		),
	)
	return err
}

// CommitTx commits specified active transaction.
func (tx *Transaction) CommitTx(ctx context.Context, opts ...CommitTransactionOption) (result *Result, err error) {
	tx.s.c.traceCommitTransactionStart(ctx, tx)
	defer func() {
		tx.s.c.traceCommitTransactionDone(ctx, tx, err)
	}()
	res := new(Ydb_Table.CommitTransactionResult)
	req := &Ydb_Table.CommitTransactionRequest{
		SessionId: tx.s.ID,
		TxId:      tx.id,
	}
	for _, opt := range opts {
		opt((*commitTransactionDesc)(req))
	}
	_, err = tx.s.c.Driver.Call(
		ydb.WithEndpointInfo(
			ctx,
			tx.s.endpointInfo,
		),
		internal.Wrap(
			Ydb_Table_V1.CommitTransaction,
			req,
			res,
		),
	)
	return &Result{stats: res.QueryStats}, err
}

// Rollback performs a rollback of the specified active transaction.
func (tx *Transaction) Rollback(ctx context.Context) (err error) {
	tx.s.c.traceRollbackTransactionStart(ctx, tx)
	defer func() {
		tx.s.c.traceRollbackTransactionDone(ctx, tx, err)
	}()
	req := Ydb_Table.RollbackTransactionRequest{
		SessionId: tx.s.ID,
		TxId:      tx.id,
	}
	_, err = tx.s.c.Driver.Call(
		ydb.WithEndpointInfo(
			ctx,
			tx.s.endpointInfo,
		),
		internal.Wrap(
			Ydb_Table_V1.RollbackTransaction,
			&req,
			nil,
		),
	)
	return err
}

func (tx *Transaction) txc() *TransactionControl {
	if tx.c == nil {
		tx.c = TxControl(WithTx(tx))
	}
	return tx.c
}

func (t *Client) traceCreateSessionStart(ctx context.Context) {
	x := CreateSessionStartInfo{
		Context: ctx,
	}
	if a := t.Trace.CreateSessionStart; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).CreateSessionStart; b != nil {
		b(x)
	}
}
func (t *Client) traceCreateSessionDone(ctx context.Context, s *Session, latency time.Duration, err error) {
	x := CreateSessionDoneInfo{
		Context: ctx,
		Session: s,
		Latency: latency,
		Error:   err,
	}
	if s != nil && s.endpointInfo != nil {
		x.Endpoint = s.endpointInfo.Address()
	}
	if a := t.Trace.CreateSessionDone; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).CreateSessionDone; b != nil {
		b(x)
	}
}
func (t *Client) traceKeepAliveStart(ctx context.Context, s *Session) {
	x := KeepAliveStartInfo{
		Context: ctx,
		Session: s,
	}
	if a := t.Trace.KeepAliveStart; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).KeepAliveStart; b != nil {
		b(x)
	}
}
func (t *Client) traceKeepAliveDone(ctx context.Context, s *Session, info SessionInfo, err error) {
	x := KeepAliveDoneInfo{
		Context:     ctx,
		Session:     s,
		SessionInfo: info,
		Error:       err,
	}
	if a := t.Trace.KeepAliveDone; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).KeepAliveDone; b != nil {
		b(x)
	}
}
func (t *Client) traceDeleteSessionStart(ctx context.Context, s *Session) {
	x := DeleteSessionStartInfo{
		Context: ctx,
		Session: s,
	}
	if a := t.Trace.DeleteSessionStart; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).DeleteSessionStart; b != nil {
		b(x)
	}
}
func (t *Client) traceDeleteSessionDone(ctx context.Context, s *Session, latency time.Duration, err error) {
	x := DeleteSessionDoneInfo{
		Context: ctx,
		Session: s,
		Latency: latency,
		Error:   err,
	}
	if a := t.Trace.DeleteSessionDone; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).DeleteSessionDone; b != nil {
		b(x)
	}
}
func (t *Client) tracePrepareDataQueryStart(ctx context.Context, s *Session, query string) {
	x := PrepareDataQueryStartInfo{
		Context: ctx,
		Session: s,
		Query:   query,
	}
	if a := t.Trace.PrepareDataQueryStart; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).PrepareDataQueryStart; b != nil {
		b(x)
	}
}

func (t *Client) tracePrepareDataQueryDone(ctx context.Context, s *Session, query string, q *DataQuery, cached bool, err error) {
	x := PrepareDataQueryDoneInfo{
		Context: ctx,
		Session: s,
		Query:   query,
		Result:  q,
		Cached:  cached,
		Error:   err,
	}
	if a := t.Trace.PrepareDataQueryDone; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).PrepareDataQueryDone; b != nil {
		b(x)
	}
}
func (t *Client) traceExecuteDataQueryStart(ctx context.Context, s *Session, tx *TransactionControl, query *DataQuery, params *QueryParameters) {
	x := ExecuteDataQueryStartInfo{
		Context:    ctx,
		Session:    s,
		TxID:       tx.id(),
		Query:      query,
		Parameters: params,
	}
	if a := t.Trace.ExecuteDataQueryStart; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).ExecuteDataQueryStart; b != nil {
		b(x)
	}
}
func (t *Client) traceExecuteDataQueryDone(
	ctx context.Context, s *Session, _ *TransactionControl, query *DataQuery,
	params *QueryParameters, prepared bool, txr *Transaction, r *Result, err error,
) {
	x := ExecuteDataQueryDoneInfo{
		Context:    ctx,
		Session:    s,
		Query:      query,
		Parameters: params,
		Prepared:   prepared,
		Result:     r,
		Error:      err,
	}
	if txr != nil {
		x.TxID = txr.id
	}
	if a := t.Trace.ExecuteDataQueryDone; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).ExecuteDataQueryDone; b != nil {
		b(x)
	}
}
func (t *Client) traceStreamReadTableStart(ctx context.Context, s *Session) {
	x := StreamReadTableStartInfo{
		Context: ctx,
		Session: s,
	}
	if a := t.Trace.StreamReadTableStart; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).StreamReadTableStart; b != nil {
		b(x)
	}
}
func (t *Client) traceStreamReadTableDone(
	ctx context.Context, s *Session, r *Result, err error,
) {
	x := StreamReadTableDoneInfo{
		Context: ctx,
		Session: s,
		Result:  r,
		Error:   err,
	}
	if a := t.Trace.StreamReadTableDone; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).StreamReadTableDone; b != nil {
		b(x)
	}
}
func (t *Client) traceStreamExecuteScanQueryStart(ctx context.Context, s *Session, query *DataQuery, params *QueryParameters) {
	x := StreamExecuteScanQueryStartInfo{
		Context:    ctx,
		Session:    s,
		Query:      query,
		Parameters: params,
	}
	if a := t.Trace.StreamExecuteScanQueryStart; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).StreamExecuteScanQueryStart; b != nil {
		b(x)
	}
}
func (t *Client) traceStreamExecuteScanQueryDone(
	ctx context.Context, s *Session, query *DataQuery,
	params *QueryParameters, r *Result, err error,
) {
	x := StreamExecuteScanQueryDoneInfo{
		Context:    ctx,
		Session:    s,
		Query:      query,
		Parameters: params,
		Result:     r,
		Error:      err,
	}
	if a := t.Trace.StreamExecuteScanQueryDone; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).StreamExecuteScanQueryDone; b != nil {
		b(x)
	}
}

func (t *Client) traceBeginTransactionStart(ctx context.Context, s *Session) {
	x := BeginTransactionStartInfo{
		Context: ctx,
		Session: s,
	}
	if a := t.Trace.BeginTransactionStart; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).BeginTransactionStart; b != nil {
		b(x)
	}
}
func (t *Client) traceBeginTransactionDone(ctx context.Context, s *Session, tx *Transaction, err error) {
	x := BeginTransactionDoneInfo{
		Context: ctx,
		Session: s,
		Error:   err,
	}
	if tx != nil {
		x.TxID = tx.id
	}
	if a := t.Trace.BeginTransactionDone; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).BeginTransactionDone; b != nil {
		b(x)
	}
}
func (t *Client) traceCommitTransactionStart(ctx context.Context, tx *Transaction) {
	x := CommitTransactionStartInfo{
		Context: ctx,
	}
	if tx != nil {
		x.Session = tx.s
		x.TxID = tx.id
	}
	if a := t.Trace.CommitTransactionStart; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).CommitTransactionStart; b != nil {
		b(x)
	}
}
func (t *Client) traceCommitTransactionDone(ctx context.Context, tx *Transaction, err error) {
	x := CommitTransactionDoneInfo{
		Context: ctx,
		Error:   err,
	}
	if tx != nil {
		x.Session = tx.s
		x.TxID = tx.id
	}
	if a := t.Trace.CommitTransactionDone; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).CommitTransactionDone; b != nil {
		b(x)
	}
}
func (t *Client) traceRollbackTransactionStart(ctx context.Context, tx *Transaction) {
	x := RollbackTransactionStartInfo{
		Context: ctx,
	}
	if tx != nil {
		x.Session = tx.s
		x.TxID = tx.id
	}
	if a := t.Trace.RollbackTransactionStart; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).RollbackTransactionStart; b != nil {
		b(x)
	}
}
func (t *Client) traceRollbackTransactionDone(ctx context.Context, tx *Transaction, err error) {
	x := RollbackTransactionDoneInfo{
		Context: ctx,
		Error:   err,
	}
	if tx != nil {
		x.Session = tx.s
		x.TxID = tx.id
	}
	if a := t.Trace.RollbackTransactionDone; a != nil {
		a(x)
	}
	if b := ContextClientTrace(ctx).RollbackTransactionDone; b != nil {
		b(x)
	}
}

type DataQuery struct {
	query    Ydb_Table.Query
	queryID  Ydb_Table.Query_Id
	queryYQL Ydb_Table.Query_YqlText
}

func (q *DataQuery) String() string {
	var emptyID Ydb_Table.Query_Id
	if q.queryID == emptyID {
		return q.queryYQL.YqlText
	}
	return q.queryID.Id
}

func (q *DataQuery) ID() string {
	return q.queryID.Id
}

func (q *DataQuery) YQL() string {
	return q.queryYQL.YqlText
}

func (q *DataQuery) initFromText(s string) {
	q.queryID = Ydb_Table.Query_Id{} // Reset id field.
	q.queryYQL.YqlText = s
	q.query.Query = &q.queryYQL
}

func (q *DataQuery) initPrepared(id string) {
	q.queryYQL = Ydb_Table.Query_YqlText{} // Reset yql field.
	q.queryID.Id = id
	q.query.Query = &q.queryID
}

func (q *DataQuery) initPreparedText(s, id string) {
	q.queryYQL = Ydb_Table.Query_YqlText{} // Reset yql field.
	q.queryYQL.YqlText = s

	q.queryID = Ydb_Table.Query_Id{} // Reset id field.
	q.queryID.Id = id

	q.query.Query = &q.queryID // Prefer preared query.
}

type QueryParameters struct {
	m queryParams
}

func (q *QueryParameters) params() queryParams {
	if q == nil {
		return nil
	}
	return q.m
}

func (q *QueryParameters) Each(it func(name string, value ydb.Value)) {
	if q == nil {
		return
	}
	for key, value := range q.m {
		it(key, internal.ValueFromYDB(
			value.Type,
			value.Value,
		))
	}
}

func (q *QueryParameters) String() string {
	var buf bytes.Buffer
	buf.WriteByte('(')
	q.Each(func(name string, value ydb.Value) {
		buf.WriteString("((")
		buf.WriteString(name)
		buf.WriteByte(')')
		buf.WriteByte('(')
		internal.WriteValueStringTo(&buf, value)
		buf.WriteString("))")
	})
	buf.WriteByte(')')
	return buf.String()
}

type queryParams map[string]*Ydb.TypedValue

type ParameterOption func(queryParams)

func NewQueryParameters(opts ...ParameterOption) *QueryParameters {
	q := &QueryParameters{
		m: make(queryParams, len(opts)),
	}
	q.Add(opts...)
	return q
}

func (q *QueryParameters) Add(opts ...ParameterOption) {
	for _, opt := range opts {
		opt(q.m)
	}
}

func ValueParam(name string, v ydb.Value) ParameterOption {
	return func(q queryParams) {
		q[name] = internal.ValueToYDB(v)
	}
}
