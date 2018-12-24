package cmd

import (
	"os"
	"time"

	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an interval manually",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.MustOpen(FlagDatabase, FlagDebug)
		defer db.Close()

		start, end := FlagDateStart, FlagDateEnd
		if start.IsZero() {
			start = time.Now()
		}

		i := interval.Interval{
			Tag:   FlagTag,
			Start: start,
			End:   end,
		}

		var count int

		if i.IsOpen() {
			db.Model(&interval.Interval{}).
				Where("tag = ? AND end = ?", i.Tag, time.Time{}).
				Count(&count)

			if count > 0 {
				cmd.Printf("there is already an open interval for tag %q\n", i.Tag)
				os.Exit(1)
			}
		} else {
			var intervals []interval.Interval

			db.
				Where(
					"(start <= $1 AND end >= $1) OR (start <= $2 AND end >= $2)",
					i.Start,
					i.End,
				).
				Find(&intervals)

			if len(intervals) > 0 {
				interval.RenderTable(cmd.OutOrStdout(), intervals...)
				cmd.Printf("\nthere already are intervals for tag %q which overlap with the specified interval\n", i.Tag)
				os.Exit(1)
			}
		}

		db.Save(&i)

		interval.RenderTable(cmd.OutOrStdout(), i)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
