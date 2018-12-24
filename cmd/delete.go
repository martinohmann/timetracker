package cmd

import (
	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "Delete an interval",
	Run:     del,
}

func init() {
	initializeIdFlag(deleteCmd)
	deleteCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(deleteCmd)
}

func del(cmd *cobra.Command, args []string) {
	database.Init(viper.GetString("database"))
	defer database.Close()

	i, err := database.FindIntervalByID(id)
	exitOnError(err)

	exitOnError(database.DeleteInterval(i))

	table.Render(cmd.OutOrStdout(), i)

	cmd.Printf("interval with ID %d deleted\n", id)
}
