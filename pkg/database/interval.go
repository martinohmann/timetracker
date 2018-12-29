package database

import (
	"time"

	"github.com/martinohmann/timetracker/pkg/interval"
)

// FindIntervalByID finds the interval with given id
func FindIntervalByID(id int) (i interval.Interval, err error) {
	return i, db.First(&i, id).Error
}

// DeleteInterval delete interval i
func DeleteInterval(i interval.Interval) error {
	return db.Delete(i).Error
}

// SaveInterval saves interval i
func SaveInterval(i *interval.Interval) error {
	return db.Save(i).Error
}

// FindOverlappingIntervals finds intervals that are overlapping with i
func FindOverlappingIntervals(i interval.Interval) (is []interval.Interval, err error) {
	err = db.
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
func FindOpenIntervalsForTag(tag string) (is []interval.Interval, err error) {
	err = db.Where("tag = ?", tag).
		Where("end = ?", time.Time{}).
		Find(&is).Error

	return is, excludeRecordNotFoundError(err)
}

// FindLastOpenIntervalForTag finds the last open interval for tag
func FindLastOpenIntervalForTag(tag string) (i interval.Interval, err error) {
	return i, db.Where("tag = ?", tag).
		Where("end = ?", time.Time{}).
		Last(&i).Error
}

// FindIntervalsByCriteria finds intervals matching the criteria
func FindIntervalsByCriteria(c interval.Interval) (is []interval.Interval, err error) {
	stmt := db
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
