package db

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Open() (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite3", "database/test.sqlite3")
	if err != nil {
		return
	}

	return
}
