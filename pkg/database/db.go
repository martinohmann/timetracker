package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/martinohmann/timetracker/pkg/tracking"
)

func MustOpen(filename string) *gorm.DB {
	db, err := gorm.Open("sqlite3", filename)
	if err != nil {
		panic(err)
	}

	return db.AutoMigrate(
		&tracking.Tracking{},
		&interval.Interval{},
	).Set("gorm:auto_preload", true)
}
