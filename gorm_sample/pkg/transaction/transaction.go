package transaction

import (
	"context"
)

type Transaction interface {
	Begin(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

func DoTransaction[T any](t Transaction, ctx context.Context, fn func(ctx context.Context) (*T, error)) (*T, error) {
	ctx, err := t.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			t.Rollback(ctx)
			panic(r)
		}
	}()
	res, err := fn(ctx)
	if err != nil {
		t.Rollback(ctx)
		return nil, err
	} else {
		t.Commit(ctx)
		return res, nil
	}
}
