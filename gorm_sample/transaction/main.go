package main

import (
	"fmt"
	"gorm_sample/pkg/models"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// ref: https://github.com/go-gorm/postgres
	db, _ := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=postgres dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Tokyo",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	uuid, _ := uuid.NewV7()

	var animals []models.Animal

	// This shows before data
	fmt.Println("Before Data")
	_ = db.Find(&animals)
	for _, animal := range animals {
		fmt.Printf("ID: %s, Name: %s\n", animal.ID, animal.Name)
	}

	tx := db.Begin()

	// This will be rolled back
	tx.Create(&models.Animal{ID: uuid.String(), Name: "Pero1"})
	// This will not be rolled back
	db.Create(&models.Animal{ID: uuid.String(), Name: "Pero2"})

	_ = tx.Find(&animals)
	fmt.Println("In Transaction")
	for _, animal := range animals {
		fmt.Printf("ID: %s, Name: %s\n", animal.ID, animal.Name)
	}
	tx.Rollback()

	// This shows remaining data
	fmt.Println("After Rollback")
	_ = db.Find(&animals)
	for _, animal := range animals {
		fmt.Printf("ID: %s, Name: %s\n", animal.ID, animal.Name)
	}
}
