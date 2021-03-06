package database

import (
	"github.com/martinohmann/timetracker/pkg/interval"
)

// Datastore defines the interface for a datastore
type Datastore interface {
	// FindIntervalByID finds the interval with given id
	FindIntervalByID(int) (interval.Interval, error)

	// DeleteInterval delete interval i
	DeleteInterval(interval.Interval) error

	// SaveInterval saves interval i
	SaveInterval(*interval.Interval) error

	// FindOverlappingIntervals finds intervals that are overlapping with i
	FindOverlappingIntervals(interval.Interval) ([]interval.Interval, error)

	// FindOpenIntervalsForTag finds open intervals for tag
	FindOpenIntervalsForTag(string) ([]interval.Interval, error)

	// FindLastOpenIntervalForTag finds the last open interval for tag
	FindLastOpenIntervalForTag(string) (interval.Interval, error)

	// FindIntervalsByCriteria finds intervals matching the criteria
	FindIntervalsByCriteria(interval.Interval) ([]interval.Interval, error)

	// Close closes the database
	Close() error
}

// FactoryFunc defines a function that creates a datastore
type FactoryFunc func(args ...interface{}) (Datastore, error)

var (
	datastore Datastore
	factory   FactoryFunc
)

func init() {
	SetFactory(SqliteFactory)
}

// SetFactory sets the factory for creating the datastore
func SetFactory(fn FactoryFunc) {
	factory = fn
}

// Init opens and initializes the database
func Init(args ...interface{}) {
	var err error

	if datastore, err = factory(args...); err != nil {
		panic(err)
	}
}

// FindIntervalByID finds the interval with given id
func FindIntervalByID(id int) (interval.Interval, error) {
	return datastore.FindIntervalByID(id)
}

// DeleteInterval delete interval i
func DeleteInterval(i interval.Interval) error {
	return datastore.DeleteInterval(i)
}

// SaveInterval saves interval i
func SaveInterval(i *interval.Interval) error {
	return datastore.SaveInterval(i)
}

// FindOverlappingIntervals finds intervals that are overlapping with i
func FindOverlappingIntervals(i interval.Interval) ([]interval.Interval, error) {
	return datastore.FindOverlappingIntervals(i)
}

// FindOpenIntervalsForTag finds open intervals for tag
func FindOpenIntervalsForTag(tag string) ([]interval.Interval, error) {
	return datastore.FindOpenIntervalsForTag(tag)
}

// FindLastOpenIntervalForTag finds the last open interval for tag
func FindLastOpenIntervalForTag(tag string) (interval.Interval, error) {
	return datastore.FindLastOpenIntervalForTag(tag)
}

// FindIntervalsByCriteria finds intervals matching the criteria
func FindIntervalsByCriteria(c interval.Interval) ([]interval.Interval, error) {
	return datastore.FindIntervalsByCriteria(c)
}

// Close closes the database
func Close() error {
	return datastore.Close()
}
