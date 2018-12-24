package cmd

import (
	"os"
	"time"

	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/dateutil"
	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show all intervals",
	Run: func(cmd *cobra.Command, args []string) {
		show(cmd, FlagDateStart, FlagDateEnd, FlagTag)
	},
}

var showYearCmd = &cobra.Command{
	Use:   "year",
	Short: "Show year's intervals",
	Run: func(cmd *cobra.Command, args []string) {
		start := dateutil.BeginOfDay(FlagYear, time.January, 1)
		end := start.AddDate(1, 0, 0)

		show(cmd, start, end, FlagTag)
	},
}

var showMonthCmd = &cobra.Command{
	Use:   "month",
	Short: "Show month's intervals",
	Run: func(cmd *cobra.Command, args []string) {
		start := dateutil.BeginOfDay(FlagYear, time.Month(FlagMonth), 1)
		end := start.AddDate(0, 1, 0)

		show(cmd, start, end, FlagTag)
	},
}

var showWeekCmd = &cobra.Command{
	Use:   "week",
	Short: "Show week's intervals",
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		FlagDate, err = dateutil.ParseDate(FlagDateString, time.Now())
		return
	},
	Run: func(cmd *cobra.Command, args []string) {
		date := FlagDate

		for date.Weekday() != time.Monday {
			date = date.AddDate(0, 0, -1)
		}

		start := dateutil.BeginOfDay(date.Year(), date.Month(), date.Day())
		end := start.AddDate(0, 0, 7)

		show(cmd, start, end, FlagTag)
	},
}

var showDateCmd = &cobra.Command{
	Use:     "date",
	Aliases: []string{"day"},
	Short:   "Show date's intervals",
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		FlagDate, err = dateutil.ParseDate(FlagDateString, time.Now())
		return
	},
	Run: func(cmd *cobra.Command, args []string) {
		date := FlagDate
		start := dateutil.BeginOfDay(date.Year(), date.Month(), date.Day())
		end := start.AddDate(0, 0, 1)

		show(cmd, start, end, FlagTag)
	},
}

func show(cmd *cobra.Command, start, end time.Time, tag string) {
	db := database.MustOpen(FlagDatabase, FlagDebug)
	defer db.Close()

	stmt := db

	if !start.IsZero() {
		stmt = stmt.Where("start >= ?", start)
	}

	if !end.IsZero() {
		stmt = stmt.Where("start < ? AND end < ?", end, end)
	}

	if tag != "" {
		stmt = stmt.Where("tag = ?", tag)
	}

	var intervals []interval.Interval
	if err := stmt.Find(&intervals).Error; err != nil {
		cmd.Println(err)
		os.Exit(1)
	}

	interval.RenderTable(cmd.OutOrStdout(), intervals...)
}

func init() {
	now := time.Now()

	showYearCmd.Flags().IntVarP(&FlagYear, "year", "", now.Year(), "filter year")
	showMonthCmd.Flags().AddFlagSet(showYearCmd.Flags())
	showMonthCmd.Flags().IntVarP(&FlagMonth, "month", "", int(now.Month()), "filter month")
	showDateCmd.Flags().StringVarP(&FlagDateString, "date", "", "", "filter date")
	showWeekCmd.Flags().AddFlagSet(showDateCmd.Flags())

	showCmd.AddCommand(showYearCmd)
	showCmd.AddCommand(showMonthCmd)
	showCmd.AddCommand(showWeekCmd)
	showCmd.AddCommand(showDateCmd)

	rootCmd.AddCommand(showCmd)
}
