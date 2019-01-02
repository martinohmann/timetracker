package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/spf13/viper"
	tilde "gopkg.in/mattes/go-expand-tilde.v1"
)

// SqliteFactory creates a new sqlite datastore
var SqliteFactory = FactoryFunc(func(args ...interface{}) (Datastore, error) {
	if len(args) == 0 {
		return nil, errors.New("invalid data source")
	}

	switch filename := args[0].(type) {
	case string:
		return newSqlite(filename)
	default:
		return nil, fmt.Errorf("invalid database source: expected string, got %v", filename)
	}
})

type sqliteDatastore struct {
	db *gorm.DB
}

// newSqlite creates a new sqlite datastore and connects to it
func newSqlite(filename string) (Datastore, error) {
	var err error
	var db *gorm.DB

	if filename, err = tilde.Expand(filename); err != nil {
		return nil, err
	}

	if db, err = gorm.Open("sqlite3", filename); err != nil {
		return nil, err
	}

	db = db.LogMode(viper.GetBool("debug")).
		AutoMigrate(&interval.Interval{})

	return &sqliteDatastore{db: db}, nil
}

// FindIntervalByID finds the interval with given id
func (d *sqliteDatastore) FindIntervalByID(id int) (i interval.Interval, err error) {
	return i, d.db.First(&i, id).Error
}

// DeleteInterval delete interval i
func (d *sqliteDatastore) DeleteInterval(i interval.Interval) error {
	return d.db.Delete(i).Error
}

// SaveInterval saves interval i
func (d *sqliteDatastore) SaveInterval(i *interval.Interval) error {
	return d.db.Save(i).Error
}

// FindOverlappingIntervals finds intervals that are overlapping with i
func (d *sqliteDatastore) FindOverlappingIntervals(i interval.Interval) (is []interval.Interval, err error) {
	err = d.db.
		Where(
			`(start BETWEEN $1 AND $2)
				OR (end BETWEEN $1 AND $2)
				OR (start <= $1 AND end >= $2)
				OR (start <= $2 AND end = $3)`,
			i.Start,
			i.End,
			time.Time{},
		).
		Where("tag = ?", i.Tag).
		Where("id != ?", i.ID).
		Find(&is).Error

	return is, excludeRecordNotFoundError(err)
}

// FindOpenIntervalsForTag finds open intervals for tag
func (d *sqliteDatastore) FindOpenIntervalsForTag(tag string) (is []interval.Interval, err error) {
	err = d.db.Where("tag = ?", tag).
		Where("end = ?", time.Time{}).
		Find(&is).Error

	return is, excludeRecordNotFoundError(err)
}

// FindLastOpenIntervalForTag finds the last open interval for tag
func (d *sqliteDatastore) FindLastOpenIntervalForTag(tag string) (i interval.Interval, err error) {
	return i, d.db.Where("tag = ?", tag).
		Where("end = ?", time.Time{}).
		Last(&i).Error
}

// FindIntervalsByCriteria finds intervals matching the criteria
func (d *sqliteDatastore) FindIntervalsByCriteria(c interval.Interval) (is []interval.Interval, err error) {
	stmt := d.db
	if !c.Start.IsZero() {
		stmt = stmt.Where("start >= ?", c.Start)
	}

	if !c.End.IsZero() {
		stmt = stmt.Where("start < ? AND end < ?", c.End, c.End)
	}

	if c.Tag != "" {
		stmt = stmt.Where("tag = ?", c.Tag)
	}

	err = stmt.Find(&is).Error

	return is, excludeRecordNotFoundError(err)
}

// Close closes the database
func (d *sqliteDatastore) Close() error {
	return d.db.Close()
}
