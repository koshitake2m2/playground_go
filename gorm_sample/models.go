package main

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Article struct {
	ID          uint `gorm:"primarykey"`
	Author      string
	Title       string
	Description string
}

type User struct {
	ID   uint `gorm:"primarykey"`
	Name string
	Age  int
}
