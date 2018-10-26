package db

import (
	"github.com/jinzhu/gorm"
)

// Open opens database and returns DB object.
func Open() (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite3", "database/test.sqlite3")
	if err != nil {
		return
	}

	return
}
