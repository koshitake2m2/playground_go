package transactionfn

import (
	"context"

	"gorm.io/gorm"
)

type Transaction interface {
	Do(ctx context.Context, fn func(ctx context.Context) error) error
}

type contextKey string

const txKey contextKey = "tx"

type GormTransaction struct {
	db *gorm.DB
}

// If you want to get the fn result, you should use outer variable to store the result.
func (g *GormTransaction) Do(ctx context.Context, fn func(ctx context.Context) error) error {
	tx := g.db.WithContext(ctx).Begin()
	ctx = context.WithValue(ctx, txKey, tx)

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()
	err := fn(ctx)
	if err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
		return nil
	}
}
