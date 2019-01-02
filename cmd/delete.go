package cmd

import (
	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	Use:     "delete [id]",
	Aliases: []string{"d", "del", "remove", "rm"},
	Short:   "Delete an interval",
	PreRunE: parseIdArg,
	Run:     del,
}

func init() {
	initializeIdFlag(deleteCmd)
	rootCmd.AddCommand(deleteCmd)
}

func del(cmd *cobra.Command, args []string) {
	database.Init(viper.GetString("database"))
	defer database.Close()

	i, err := database.FindIntervalByID(id)
	exitOnError(cmd, err)

	exitOnError(cmd, database.DeleteInterval(i))

	table.Render(cmd.OutOrStdout(), i)

	cmd.Printf("interval with ID %d deleted\n", i.ID)
}
