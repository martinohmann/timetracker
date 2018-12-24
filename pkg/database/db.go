package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/martinohmann/timetracker/pkg/interval"
)

func MustOpen(filename string, debug bool) *gorm.DB {
	db, err := gorm.Open("sqlite3", filename)
	if err != nil {
		panic(err)
	}

	return db.LogMode(debug).AutoMigrate(&interval.Interval{})
}
