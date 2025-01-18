package transaction

import (
	"context"

	"gorm.io/gorm"
)

type contextKey string

const txKey contextKey = "tx"

func GetTx(ctx context.Context) *gorm.DB {
	return ctx.Value(txKey).(*gorm.DB)
}

type GormTransaction struct {
	db *gorm.DB
}

func (g *GormTransaction) Begin(ctx context.Context) (context.Context, error) {
	tx := g.db.WithContext(ctx).Begin()
	ctx = context.WithValue(ctx, txKey, tx)
	return ctx, nil
}

func (g *GormTransaction) Commit(ctx context.Context) error {
	tx := GetTx(ctx)
	tx.Commit()
	return nil
}

func (g *GormTransaction) Rollback(ctx context.Context) error {
	tx := GetTx(ctx)
	tx.Rollback()
	return nil
}
