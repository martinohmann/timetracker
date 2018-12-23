package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/martinohmann/timetracker/pkg/tracking"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start time tracking",
	Long:  `Long description`,
	Run: func(cmd *cobra.Command, args []string) {
		db := database.MustOpen(DatabaseFile)
		defer db.Close()

		var count int

		db.Model(&tracking.Tracking{}).
			Where("tag = ? AND finished = ?", FlagTag, false).
			Count(&count)

		if count > 0 {
			fmt.Printf("there is already a tracking running for tag %q\n", FlagTag)
			os.Exit(1)
		}

		t := tracking.Tracking{
			Tag: FlagTag,
			Interval: interval.Interval{
				From: time.Now(),
			},
		}

		db.Save(&t)

		tracking.RenderTable(os.Stdout, t)
	},
}

func init() {
	startCmd.Flags().StringVarP(&FlagTag, "tag", "t", tracking.DefaultTag, "Tracking tag to use")
	rootCmd.AddCommand(startCmd)
}
