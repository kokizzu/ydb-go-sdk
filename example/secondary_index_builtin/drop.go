package main

import (
	"context"
	"path"

	"github.com/yandex-cloud/ydb-go-sdk/table"
)

func doDrop(
	ctx context.Context,
	sp *table.SessionPool,
	prefix string,
	args ...string,
) error {
	for _, name := range []string{"series", "users"} {
		err := table.Retry(ctx, sp,
			table.OperationFunc(func(ctx context.Context, s *table.Session) error {
				return s.DropTable(ctx, path.Join(prefix, name))
			}),
		)
		if err != nil {
			return err
		}
	}
	return nil
}
