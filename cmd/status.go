package cmd

import (
	"fmt"
	"os"

	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/tracking"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show time tracking status",
	Long:  `Long description`,
	Run: func(cmd *cobra.Command, args []string) {
		db := database.MustOpen(FlagDatabase)
		defer db.Close()

		stmt := db.Where("finished = ?", false)

		if FlagTag != "" {
			stmt = stmt.Where("tag = ?", FlagTag)
		}

		var trackings []tracking.Tracking
		if err := stmt.Find(&trackings).Error; err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tracking.RenderTable(os.Stdout, trackings...)
	},
}

func init() {
	statusCmd.Flags().StringVarP(&FlagTag, "tag", "t", "", "Tracking tag to use")
	rootCmd.AddCommand(statusCmd)
}
