package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

var db *gorm.DB

func DB() *gorm.DB {
	if db == nil {

		var err error
		// TODO: get from config
		db, err = gorm.Open("postgres", "postgres://admin:admin@postgres:5432/ketabdoozak?sslmode=disable")
		if err != nil {
			log.Panic(errors.Wrap(err, "error on connecting database"))
		}
	}

	return db
}

func Close() error {
	if db != nil {
		return db.Close()
	}

	return nil
}
