package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/spf13/viper"
	tilde "gopkg.in/mattes/go-expand-tilde.v1"
)

func MustOpen(filename string) *gorm.DB {
	filename, err := tilde.Expand(filename)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("sqlite3", filename)
	if err != nil {
		panic(err)
	}

	return db.LogMode(viper.GetBool("debug")).AutoMigrate(&interval.Interval{})
}
