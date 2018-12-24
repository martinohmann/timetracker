package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/martinohmann/timetracker/pkg/dateutil"
	"github.com/martinohmann/timetracker/pkg/version"
	"github.com/spf13/cobra"
	tilde "gopkg.in/mattes/go-expand-tilde.v1"
)

var rootCmd = &cobra.Command{
	Use:     "timetracker",
	Short:   "Track time intervals using tags",
	Version: version.Version,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		FlagDateStart, err = dateutil.ParseDate(FlagDateStartString, time.Time{})
		if err != nil {
			return
		}

		FlagDateEnd, err = dateutil.ParseDate(FlagDateEndString, time.Time{})
		return
	},
}

var (
	FlagDebug           bool
	FlagDatabase        string
	FlagDateString      string
	FlagDateStartString string
	FlagDateEndString   string
	FlagID              int
	FlagMonth           int
	FlagTag             string
	FlagYear            int

	FlagDate      time.Time
	FlagDateStart time.Time
	FlagDateEnd   time.Time
)

func init() {
	dbPath, err := tilde.Expand("~/.timetracker.db")
	if err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().StringVar(&FlagDatabase, "database", dbPath, "path to sqlite database")
	rootCmd.PersistentFlags().BoolVar(&FlagDebug, "debug", false, "enable debug output")
	rootCmd.PersistentFlags().StringVar(&FlagDateStartString, "start", "", "start date of the interval")
	rootCmd.PersistentFlags().StringVar(&FlagDateEndString, "end", "", "end date of the interval")
	rootCmd.PersistentFlags().StringVarP(&FlagTag, "tag", "t", "", "interval tag to use")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
