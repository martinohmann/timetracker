package cmd

import (
	"fmt"
	"os"

	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/tracking"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:     "del",
	Aliases: []string{"delete"},
	Short:   "Delete a tracked time",
	Long:    `Long description`,
	Run: func(cmd *cobra.Command, args []string) {
		db := database.MustOpen(DatabaseFile)
		defer db.Close()

		if FlagID == 0 {
			fmt.Println("need tracking id")
			os.Exit(1)
		}

		var t tracking.Tracking

		err := db.First(&t, FlagID).Error
		if err != nil {
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
