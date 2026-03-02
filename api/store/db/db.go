package db

import (
	"database/sql"
	"itruj/discens-vestigo/api/store"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB interface {
	Model(value any) (tx *gorm.DB)
	Rows() (*sql.Rows, error)
	Select(query any, args ...any) (tx *gorm.DB)
}

type sqliteDB struct {
	*gorm.DB
}

func (d *sqliteDB) Model(value any) (tx *gorm.DB) {
	return d.DB.Model(value)
}

func (d *sqliteDB) Rows() (*sql.Rows, error) {
	return d.DB.Rows()
}

func (d *sqliteDB) Select(query any, args ...any) (tx *gorm.DB) {
	return d.DB.Select(query, args...)
}

func open(dbName string) (*gorm.DB, error) {
	err := os.MkdirAll("/tmp", 0755)

	if err != nil {
		return nil, err
	}

	return gorm.Open(sqlite.Open(dbName), &gorm.Config{})
}

func MustOpen(dbName string) *gorm.DB {
	db, err := open(dbName)

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&store.Card{})

	if err != nil {
		panic(err)
	}

	return db
}
