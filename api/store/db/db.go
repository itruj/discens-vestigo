package db

import (
	"itruj/discens-vestigo/api/store"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func open(dbName string) (*gorm.DB, error) {

	// make the temp directory if it doesn't exist
	err := os.MkdirAll("/tmp", 0755)
	if err != nil {
		return nil, err
	}

	return gorm.Open(sqlite.Open(dbName), &gorm.Config{})
}

func MustOpen(dbName string) *gorm.DB {

	if dbName == "" {
		dbName = "discens-vestigo.db"
	}

	db, err := open(dbName)
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&store.User{}, &store.Session{})

	if err != nil {
		panic(err)
	}

	return db
}
