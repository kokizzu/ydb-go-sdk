package testutil

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/yandex-cloud/ydb-go-sdk"
	"github.com/yandex-cloud/ydb-go-sdk/api"

	"github.com/yandex-cloud/ydb-go-sdk/api/grpc/Ydb_Table_V1"
	"github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb_Table"
	"github.com/yandex-cloud/ydb-go-sdk/internal"
)

var ErrNotImplemented = errors.New("testutil: not implemented")

type MethodCode uint

func (m MethodCode) String() string {
	return codeToString[m]
}

const (
	UnknownMethod MethodCode = iota
	TableCreateSession
	TableDeleteSession
	TableKeepAlive
	TableCreateTable
	TableDropTable
	TableAlterTable
	TableCopyTable
	TableDescribeTable
	TableExplainDataQuery
	TablePrepareDataQuery
	TableExecuteDataQuery
	TableExecuteSchemeQuery
	TableBeginTransaction
	TableCommitTransaction
	TableRollbackTransaction
	TableDescribeTableOptions
	TableStreamReadTable
)

var grpcMethodToCode = map[string]MethodCode{
	Ydb_Table_V1.CreateSession:        TableCreateSession,
	Ydb_Table_V1.DeleteSession:        TableDeleteSession,
	Ydb_Table_V1.KeepAlive:            TableKeepAlive,
	Ydb_Table_V1.CreateTable:          TableCreateTable,
	Ydb_Table_V1.DropTable:            TableDropTable,
	Ydb_Table_V1.AlterTable:           TableAlterTable,
	Ydb_Table_V1.CopyTable:            TableCopyTable,
	Ydb_Table_V1.DescribeTable:        TableDescribeTable,
	Ydb_Table_V1.ExplainDataQuery:     TableExplainDataQuery,
	Ydb_Table_V1.PrepareDataQuery:     TablePrepareDataQuery,
	Ydb_Table_V1.ExecuteDataQuery:     TableExecuteDataQuery,
	Ydb_Table_V1.ExecuteSchemeQuery:   TableExecuteSchemeQuery,
	Ydb_Table_V1.BeginTransaction:     TableBeginTransaction,
	Ydb_Table_V1.CommitTransaction:    TableCommitTransaction,
	Ydb_Table_V1.RollbackTransaction:  TableRollbackTransaction,
	Ydb_Table_V1.DescribeTableOptions: TableDescribeTableOptions,
	Ydb_Table_V1.StreamReadTable:      TableStreamReadTable,
}

var codeToString = map[MethodCode]string{
	TableCreateSession:        lastSegment(Ydb_Table_V1.CreateSession),
	TableDeleteSession:        lastSegment(Ydb_Table_V1.DeleteSession),
	TableKeepAlive:            lastSegment(Ydb_Table_V1.KeepAlive),
	TableCreateTable:          lastSegment(Ydb_Table_V1.CreateTable),
	TableDropTable:            lastSegment(Ydb_Table_V1.DropTable),
	TableAlterTable:           lastSegment(Ydb_Table_V1.AlterTable),
	TableCopyTable:            lastSegment(Ydb_Table_V1.CopyTable),
	TableDescribeTable:        lastSegment(Ydb_Table_V1.DescribeTable),
	TableExplainDataQuery:     lastSegment(Ydb_Table_V1.ExplainDataQuery),
	TablePrepareDataQuery:     lastSegment(Ydb_Table_V1.PrepareDataQuery),
	TableExecuteDataQuery:     lastSegment(Ydb_Table_V1.ExecuteDataQuery),
	TableExecuteSchemeQuery:   lastSegment(Ydb_Table_V1.ExecuteSchemeQuery),
	TableBeginTransaction:     lastSegment(Ydb_Table_V1.BeginTransaction),
	TableCommitTransaction:    lastSegment(Ydb_Table_V1.CommitTransaction),
	TableRollbackTransaction:  lastSegment(Ydb_Table_V1.RollbackTransaction),
	TableDescribeTableOptions: lastSegment(Ydb_Table_V1.DescribeTableOptions),
	TableStreamReadTable:      lastSegment(Ydb_Table_V1.StreamReadTable),
}

func setField(name string, dst, value interface{}) {
	x := reflect.ValueOf(dst).Elem()
	t := x.Type()
	f, ok := t.FieldByName(name)
	if !ok {
		panic(fmt.Sprintf(
			"ydb/testutil: struct %s has no field %q",
			t, name,
		))
	}
	v := reflect.ValueOf(value)
	if f.Type.Kind() != v.Type().Kind() {
		panic(fmt.Sprintf(
			"ydb/testutil: struct %s field %q is type of %s, not %s",
			t, name, f.Type, v.Type(),
		))
	}
	x.FieldByName(f.Name).Set(v)
}

func getField(name string, src, dst interface{}) bool {
	var fn func(x reflect.Value, seg ...string) bool
	fn = func(x reflect.Value, seg ...string) bool {
		if x.Kind() == reflect.Ptr {
			x = x.Elem()
		}
		t := x.Type()
		f, ok := t.FieldByName(seg[0])
		if !ok {
			return false
		}
		fv := x.FieldByName(seg[0])
		if fv.Kind() == reflect.Ptr && fv.IsNil() {
			return false
		}
		if len(seg) > 1 {
			return fn(fv.Elem(), seg[1:]...)
		}

		v := reflect.ValueOf(dst)
		if v.Type().Kind() != reflect.Ptr {
			panic("ydb/testutil: destination value must be a pointer")
		}
		if v.Type().Elem().Kind() != fv.Type().Kind() {
			panic(fmt.Sprintf(
				"ydb/testutil: struct %s field %q is type of %s, not %s",
				t, name, f.Type, v.Type(),
			))
		}

		v.Elem().Set(fv)

		return true
	}
	return fn(reflect.ValueOf(src).Elem(), strings.Split(name, ".")...)
}

type TableCreateSessionResult struct {
	R interface{}
}

func (t TableCreateSessionResult) SetSessionID(id string) {
	setField("SessionId", t.R, id)
}

type TableKeepAliveResult struct {
	R interface{}
}

func (t TableKeepAliveResult) SetSessionStatus(ready bool) {
	var status Ydb_Table.KeepAliveResult_SessionStatus
	if ready {
		status = Ydb_Table.KeepAliveResult_SESSION_STATUS_READY
	} else {
		status = Ydb_Table.KeepAliveResult_SESSION_STATUS_BUSY
	}
	setField("SessionStatus", t.R, status)
}

type TableBeginTransactionResult struct {
	R interface{}
}

func (t TableBeginTransactionResult) SetTransactionID(id string) {
	setField("TxMeta", t.R, &Ydb_Table.TransactionMeta{
		Id: id,
	})
}

type TableExecuteDataQueryResult struct {
	R interface{}
}

func (t TableExecuteDataQueryResult) SetTransactionID(id string) {
	setField("TxMeta", t.R, &Ydb_Table.TransactionMeta{
		Id: id,
	})
}

type TableExecuteDataQueryRequest struct {
	R interface{}
}

func (t TableExecuteDataQueryRequest) SessionID() (id string) {
	getField("SessionId", t.R, &id)
	return
}

func (t TableExecuteDataQueryRequest) TransactionID() (id string, ok bool) {
	ok = getField("TxControl.TxSelector.TxId", t.R, &id)
	return
}

type TablePrepareDataQueryResult struct {
	R interface{}
}

func (t TablePrepareDataQueryResult) SetQueryID(id string) {
	setField("QueryId", t.R, id)
}

type Driver struct {
	OnCall       func(ctx context.Context, code MethodCode, req, res interface{}) error
	OnStreamRead func(ctx context.Context, code MethodCode, req, res interface{}, process func(error)) error
	OnClose      func() error
}

func (d *Driver) CallEx(ctx context.Context, op api.Operation, ex *ydb.ExtendedCallParams) (*ydb.MetaInfo, error) {
	if d.OnCall == nil {
		return nil, ErrNotImplemented
	}
	method, req, res, _ := internal.Unwrap(op)
	code := grpcMethodToCode[method]

	// NOTE: req and res may be converted to testutil inner structs, which are
	// mirrors of grpc api envelopes.
	return nil, d.OnCall(ctx, code, req, res)
}

func (d *Driver) StreamReadEx(ctx context.Context, op api.StreamOperation, ex *ydb.ExtendedCallParams) (*ydb.MetaInfo, error) {
	if d.OnStreamRead == nil {
		return nil, ErrNotImplemented
	}
	method, req, res, processor := internal.UnwrapStreamOperation(op)
	code := grpcMethodToCode[method]

	return nil, d.OnStreamRead(ctx, code, req, res, processor)
}

func (d *Driver) Call(ctx context.Context, op internal.Operation) error {
	_, err := d.CallEx(ctx, op, nil)
	return err
}

func (d *Driver) StreamRead(ctx context.Context, op internal.StreamOperation) error {
	_, err := d.StreamReadEx(ctx, op, nil)
	return err
}

func (d *Driver) Close() error {
	if d.OnClose == nil {
		return ErrNotImplemented
	}
	return d.OnClose()
}

func lastSegment(m string) string {
	s := strings.Split(m, "/")
	return s[len(s)-1]
}
