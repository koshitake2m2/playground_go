package main

import (
	"context"
	"fmt"
	"gorm_sample/pkg/models"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type contextKey string

const txKey contextKey = "tx"

func main() {

	ctx := context.Background()

	// ref: https://github.com/go-gorm/postgres
	db, _ := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=postgres dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Tokyo",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	// Commit
	tx := db.WithContext(ctx).Begin()
	ctxWithTx := context.WithValue(ctx, txKey, tx)
	Create(ctxWithTx, 1)
	tx.Commit()

	// Rollback
	tx2 := db.WithContext(ctx).Begin()
	ctxWithTx2 := context.WithValue(ctx, txKey, tx2)
	Create(ctxWithTx2, 2)
	tx2.Rollback()

}

func Create(ctx context.Context, n int64) error {
	uuid, _ := uuid.NewV7()

	tx := ctx.Value(txKey).(*gorm.DB)
	tx.Create(&models.Animal{ID: uuid.String(), Name: fmt.Sprintf("Pero in Context %d", n)})

	return nil
}
