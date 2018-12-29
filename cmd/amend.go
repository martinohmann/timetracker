package cmd

import (
	"os"

	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var amendCmd = &cobra.Command{
	Use:     "amend [id]",
	Short:   "Amend an interval",
	PreRunE: preRunE(parseIdArg, parseDateRange),
	Run:     amend,
}

func init() {
	initializeDateRangeFlags(amendCmd)
	initializeTagFlag(amendCmd)
	rootCmd.AddCommand(amendCmd)
}

func amend(cmd *cobra.Command, args []string) {
	database.Init(viper.GetString("database"))
	defer database.Close()

	i, err := database.FindIntervalByID(id)
	exitOnError(err)

	if !startDate.IsZero() {
		i.Start = startDate
	}

	if !endDate.IsZero() {
		i.End = endDate
	}

	if tag != "" {
		i.Tag = tag
	}

	intervals, err := database.FindOverlappingIntervals(i)
	exitOnError(err)

	if len(intervals) > 0 {
		table.Render(cmd.OutOrStdout(), intervals...)
		cmd.Printf("there already are intervals for tag %q which overlap with the specified interval\n", i.Tag)
		os.Exit(1)
	}

	exitOnError(database.SaveInterval(&i))

	table.Render(cmd.OutOrStdout(), i)

	cmd.Printf("interval with ID %d updated\n", i.ID)
}
