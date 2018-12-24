package cmd

import (
	"os"
	"time"

	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start an interval",
	Run:   start,
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func start(cmd *cobra.Command, args []string) {
	db := database.MustOpen(viper.GetString("database"))
	defer db.Close()

	var count int

	db.Model(&interval.Interval{}).
		Where("tag = ?", tag).
		Where("end = ?", time.Time{}).
		Count(&count)

	if count > 0 {
		cmd.Printf("there is already an open interval for tag %q\n", tag)
		os.Exit(1)
	}

	i := interval.Interval{
		Tag:   tag,
		Start: time.Now(),
	}

	db.Save(&i)

	interval.RenderTable(cmd.OutOrStdout(), i)
}
