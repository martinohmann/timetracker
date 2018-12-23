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
	Use:     "ttc",
	Short:   "Simple time tracker",
	Long:    `Long description`,
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
	FlagID              int
	FlagTag             string
	FlagDatabase        string
	FlagDateStartString string
	FlagDateEndString   string
	FlagDateString      string
	FlagYear            int
	FlagMonth           int

	FlagDateStart time.Time
	FlagDateEnd   time.Time
	FlagDate      time.Time
)

func init() {
	dbPath, err := tilde.Expand("~/.timetracker.db")
	if err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().StringVarP(&FlagDatabase, "database", "", dbPath, "Path to sqlite database")
	rootCmd.PersistentFlags().StringVarP(&FlagDateStartString, "start", "", "", "Start date for tracking")
	rootCmd.PersistentFlags().StringVarP(&FlagDateEndString, "end", "", "", "End date for tracking")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
