package cmd

import (
	"fmt"
	"os"

	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/tracking"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "Delete a tracked time",
	Long:    `Long description`,
	Run: func(cmd *cobra.Command, args []string) {
		db := database.MustOpen(FlagDatabase)
		defer db.Close()

		var t tracking.Tracking

		if err := db.First(&t, FlagID).Error; err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		db.Delete(&t)

		tracking.RenderTable(os.Stdout, t)

		fmt.Println("deleted")
	},
}

func init() {
	delCmd.Flags().IntVarP(&FlagID, "id", "", 0, "Tracking ID")
	delCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(delCmd)
}
