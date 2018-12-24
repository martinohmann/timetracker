package cmd

import (
	"fmt"
	"os"

	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "Delete an interval",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.MustOpen(FlagDatabase, FlagDebug)
		defer db.Close()

		var i interval.Interval

		if err := db.First(&i, FlagID).Error; err != nil {
			cmd.Println(err)
			os.Exit(1)
		}

		db.Delete(&i)

		interval.RenderTable(cmd.OutOrStdout(), i)

		fmt.Println("deleted")
	},
}

func init() {
	deleteCmd.Flags().IntVarP(&FlagID, "id", "", 0, "interval ID")
	deleteCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(deleteCmd)
}
