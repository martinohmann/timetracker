package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/martinohmann/timetracker/pkg/dateutil"
	"github.com/martinohmann/timetracker/pkg/version"
	"github.com/spf13/cobra"
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
	FlagDateStartString string
	FlagDateEndString   string
	FlagDateString      string
	FlagYear            int
	FlagMonth           int

	FlagDateStart time.Time
	FlagDateEnd   time.Time
	FlagDate      time.Time
)

const DatabaseFile = "/tmp/test.db"

func init() {
	rootCmd.PersistentFlags().StringVarP(&FlagDateStartString, "from", "", "", "Start date for tracking")
	rootCmd.PersistentFlags().StringVarP(&FlagDateEndString, "to", "", "", "End date for tracking")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
