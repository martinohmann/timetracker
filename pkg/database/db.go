package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/spf13/viper"
	tilde "gopkg.in/mattes/go-expand-tilde.v1"
)

var db *gorm.DB

// Init opens and initializes the database
func Init(filename string) {
	var err error

	if filename, err = tilde.Expand(filename); err != nil {
		panic(err)
	}

	if db, err = gorm.Open("sqlite3", filename); err != nil {
		panic(err)
	}

	db = db.LogMode(viper.GetBool("debug")).
		AutoMigrate(&interval.Interval{})
}

// Close closes the database
func Close() {
	if db != nil {
		db.Close()
	}
}

func excludeRecordNotFoundError(err error) error {
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}

	return nil
}
