package cmd

import (
	"os"
	"time"

	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an interval manually",
	Run:   add,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func add(cmd *cobra.Command, args []string) {
	db := database.MustOpen(viper.GetString("database"))
	defer db.Close()

	if startDate.IsZero() {
		startDate = time.Now()
	}

	i := interval.Interval{
		Tag:   tag,
		Start: startDate,
		End:   endDate,
	}

	var count int

	if i.IsClosed() {
		var intervals []interval.Interval

		db.
			Where("tag = ?", i.Tag).
			Where(
				"(start <= ? AND end >= ?) OR (start <= ? AND end >= ?) OR (end = ?)",
				i.Start,
				i.Start,
				i.End,
				i.End,
				time.Time{},
			).
			Find(&intervals)

		if len(intervals) > 0 {
			interval.RenderTable(cmd.OutOrStdout(), intervals...)
			cmd.Printf("\nthere already are intervals for tag %q which overlap with the specified interval\n", i.Tag)
			os.Exit(1)
		}
	} else {
		db.Model(&interval.Interval{}).
			Where("tag = ? AND end = ?", i.Tag, time.Time{}).
			Count(&count)

		if count > 0 {
			cmd.Printf("there is already an open interval for tag %q\n", i.Tag)
			os.Exit(1)
		}
	}

	db.Save(&i)

	interval.RenderTable(cmd.OutOrStdout(), i)
}
