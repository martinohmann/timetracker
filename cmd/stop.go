package cmd

import (
	"os"

	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/martinohmann/timetracker/pkg/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var stopCmd = &cobra.Command{
	Use:     "stop [tag]",
	Aliases: []string{"c", "close"},
	Short:   "Stop an interval",
	PreRunE: parseTagArg,
	Run:     stop,
}

func init() {
	initializeIdFlag(stopCmd)
	initializeTagFlag(stopCmd)
	rootCmd.AddCommand(stopCmd)
}

func stop(cmd *cobra.Command, args []string) {
	database.Init(viper.GetString("database"))
	defer database.Close()

	var err error
	var i interval.Interval

	if id > 0 {
		i, err = database.FindIntervalByID(id)
	} else {
		i, err = database.FindLastOpenIntervalForTag(tag)
	}

	exitOnError(err)

	if i.IsClosed() {
		cmd.Printf("interval %d is already closed\n", i.ID)
		os.Exit(1)
	}

	i.Close()

	exitOnError(database.SaveInterval(&i))

	table.Render(cmd.OutOrStdout(), i)

	cmd.Printf("interval with ID %d closed\n", id)
}
