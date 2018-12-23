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

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add an interval manually",
	Long:  `Long description`,
	Run: func(cmd *cobra.Command, args []string) {
		db := database.MustOpen(FlagDatabase)
		defer db.Close()

		start, end := FlagDateStart, FlagDateEnd
		if start.IsZero() {
			start = time.Now()
		}

		t := tracking.Tracking{
			Tag:      FlagTag,
			Finished: !end.IsZero(),
			Interval: interval.Interval{
				Start: start,
				End:   end,
			},
		}

		var count int

		if !t.Finished {
			db.Model(&tracking.Tracking{}).
				Where("tag = ? AND finished = ?", t.Tag, false).
				Count(&count)

			if count > 0 {
				fmt.Printf("there is already a tracking running for tag %q\n", t.Tag)
				os.Exit(1)
			}
		} else {
			var trackings []tracking.Tracking

			db.Joins("JOIN intervals ON trackings.interval_id = intervals.id").
				Where(
					"(intervals.start <= ? AND intervals.end >= ?) OR "+
						"(intervals.start <= ? AND intervals.end >= ?)",
					t.Interval.Start,
					t.Interval.Start,
					t.Interval.End,
					t.Interval.End,
				).
				Find(&trackings)

			if len(trackings) > 0 {
				tracking.RenderTable(os.Stdout, trackings...)
				fmt.Printf("\nthere already are trackings for tag %q which overlaps with the specified interval\n", t.Tag)
				os.Exit(1)
			}
		}

		db.Save(&t)

		tracking.RenderTable(os.Stdout, t)
	},
}

func init() {
	addCmd.Flags().StringVarP(&FlagTag, "tag", "t", tracking.DefaultTag, "Tracking tag to use")
	rootCmd.AddCommand(addCmd)
}
