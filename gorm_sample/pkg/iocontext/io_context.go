package transaction

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type IOContext interface {
}
type IOContextHelper interface {
	Transaction(f func(io IOContext) error) error
	ReadOnly(f func(io IOContext) error) error
}

type GormIOContextHelpler struct {
}

func (i *GormIOContextHelpler) Transaction(f func(tx *gorm.DB) error) error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	tx := db.Begin()
	ferr := f(tx)
	if ferr != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return nil
}

type GormIOContext struct {
	db *gorm.DB
}

type GormIOContextHelper struct {
}

func (i *GormIOContextHelper) Transaction(f func(io IOContext) error) error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	tx := db.Begin()
	txio := GormIOContext{db: tx}
	ferr := f(txio)
	if ferr != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return ferr
}

func (i *GormIOContextHelper) ReadOnly(f func(io IOContext) error) error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	roio := GormIOContext{db: db}
	ferr := f(roio)
	return ferr
}
