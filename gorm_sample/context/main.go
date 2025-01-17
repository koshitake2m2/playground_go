package main

import (
	"context"
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
	tx := db.Begin()
	ctx = context.WithValue(ctx, txKey, tx)

	Create(ctx)

	tx.Commit()

}

func Create(ctx context.Context) error {
	uuid, _ := uuid.NewV7()

	tx := ctx.Value("tx").(*gorm.DB)
	tx.Create(&models.Animal{ID: uuid.String(), Name: "Pero in Context"})

	return nil
}
