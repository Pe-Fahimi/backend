package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

var db *gorm.DB

func DB() *gorm.DB {
	if db == nil {

		var err error
		db, err = gorm.Open("postgres", os.Getenv("DATABASE_URI"))
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
