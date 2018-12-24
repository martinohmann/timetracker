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

var addCmd = &cobra.Command{
	Use:     "add [tag]",
	Aliases: []string{"a"},
	Short:   "Add an interval manually",
	PreRunE: preRunE(parseTagArg, parseDateRange),
	Run:     add,
}

func init() {
	initializeDateRangeFlags(addCmd)
	initializeTagFlag(addCmd)
	rootCmd.AddCommand(addCmd)
}

func add(cmd *cobra.Command, args []string) {
	database.Init(viper.GetString("database"))
	defer database.Close()

	if startDate.IsZero() {
		startDate = time.Now()
	}

	i := interval.Interval{
		Tag:   tag,
		Start: startDate,
		End:   endDate,
	}

	if i.IsClosed() {
		intervals, err := database.FindOverlappingIntervals(i)
		exitOnError(err)

		if len(intervals) > 0 {
			table.Render(cmd.OutOrStdout(), intervals...)
			cmd.Printf("there already are intervals for tag %q which overlap with the specified interval\n", i.Tag)
			os.Exit(1)
		}
	} else {
		count, err := database.CountOpenIntervalsForTag(i.Tag)
		exitOnError(err)

		if count > 0 {
			cmd.Printf("there is already an open interval for tag %q\n", i.Tag)
			os.Exit(1)
		}
	}

	exitOnError(database.SaveInterval(&i))

	table.Render(cmd.OutOrStdout(), i)

	cmd.Printf("interval with tag %q added\n", tag)

}
