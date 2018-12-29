package cmd

import (
	"os"
	"time"

	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/martinohmann/timetracker/pkg/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:     "start [tag]",
	Aliases: []string{"o", "open"},
	Short:   "Start an interval",
	PreRunE: parseTagArg,
	Run:     start,
}

func init() {
	initializeTagFlag(startCmd)
	rootCmd.AddCommand(startCmd)
}

func start(cmd *cobra.Command, args []string) {
	database.Init(viper.GetString("database"))
	defer database.Close()

	intervals, err := database.FindOpenIntervalsForTag(tag)
	exitOnError(err)

	if len(intervals) > 0 {
		table.Render(cmd.OutOrStdout(), intervals...)
		cmd.Printf("there is already an open interval for tag %q\n", tag)
		os.Exit(1)
	}

	i := interval.Interval{
		Tag:   tag,
		Start: time.Now(),
	}

	exitOnError(database.SaveInterval(&i))

	table.Render(cmd.OutOrStdout(), i)

	cmd.Printf("interval with tag %q started\n", tag)
}
