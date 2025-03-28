package main

import (
	"gorm_sample/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// ref: https://github.com/go-gorm/postgres
	db, _ := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=postgres dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Tokyo",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	db.Create(&models.Article{ID: 4, Author: "author4", Title: "title4", Description: "description4"})
}
