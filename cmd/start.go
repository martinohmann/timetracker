package cmd

import (
	"os"
	"time"

	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start an interval",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.MustOpen(FlagDatabase, FlagDebug)
		defer db.Close()

		var count int

		db.Model(&interval.Interval{}).
			Where("tag = ?", FlagTag).
			Where("end = ?", time.Time{}).
			Count(&count)

		if count > 0 {
			cmd.Printf("there is already an open interval for tag %q\n", FlagTag)
			os.Exit(1)
		}

		i := interval.Interval{
			Tag:   FlagTag,
			Start: time.Now(),
		}

		db.Save(&i)

		interval.RenderTable(cmd.OutOrStdout(), i)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
