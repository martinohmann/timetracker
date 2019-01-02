package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/spf13/cobra"
)

type mockDatastore struct{}

var (
	mockFactory = database.FactoryFunc(func(...interface{}) (database.Datastore, error) {
		return &mockDatastore{}, nil
	})
)

func (d *mockDatastore) FindIntervalByID(id int) (interval.Interval, error) {
	return interval.Interval{ID: id}, nil
}

func (d *mockDatastore) FindLastOpenIntervalForTag(tag string) (interval.Interval, error) {
	return interval.Interval{Tag: tag}, nil
}

func (d *mockDatastore) DeleteInterval(i interval.Interval) error { return nil }
func (d *mockDatastore) SaveInterval(i *interval.Interval) error  { return nil }
func (d *mockDatastore) Close() error                             { return nil }

func (d *mockDatastore) FindOverlappingIntervals(i interval.Interval) (is []interval.Interval, err error) {
	return is, err
}

func (d *mockDatastore) FindOpenIntervalsForTag(tag string) (is []interval.Interval, err error) {
	return is, err
}

func (d *mockDatastore) FindIntervalsByCriteria(c interval.Interval) (is []interval.Interval, err error) {
	return is, err
}

func executeCommand(cmd *cobra.Command, args ...string) (bytes.Buffer, error) {
	cmd.SetArgs(args)

	var buf bytes.Buffer

	cmd.SetOutput(&buf)

	err := cmd.Execute()

	return buf, err
}

func TestMain(m *testing.M) {
	database.SetFactory(mockFactory)
	os.Exit(m.Run())
}
