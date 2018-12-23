package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/dateutil"
	"github.com/martinohmann/timetracker/pkg/tracking"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show time trackings",
	Long:  `Long description`,
	Run: func(cmd *cobra.Command, args []string) {
		show(FlagDateStart, FlagDateEnd, FlagTag)
	},
}

var showYearCmd = &cobra.Command{
	Use:   "year",
	Short: "Show year",
	Long:  `Long description`,
	Run: func(cmd *cobra.Command, args []string) {
		from := dateutil.BeginOfDay(FlagYear, time.January, 1)
		to := dateutil.EndOfDay(FlagYear, time.December, 31)

		show(from, to, FlagTag)
	},
}

var showMonthCmd = &cobra.Command{
	Use:   "month",
	Short: "Show month",
	Long:  `Long description`,
	Run: func(cmd *cobra.Command, args []string) {
		from := dateutil.BeginOfDay(FlagYear, time.Month(FlagMonth), 1)
		lastOfMonth := from.AddDate(0, 1, -1)
		to := dateutil.EndOfDay(lastOfMonth.Year(), lastOfMonth.Month(), lastOfMonth.Day())

		show(from, to, FlagTag)
	},
}

var showWeekCmd = &cobra.Command{
	Use:   "week",
	Short: "Show week",
	Long:  `Long description`,
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		FlagDate, err = dateutil.ParseDate(FlagDateString, time.Now())
		return
	},
	Run: func(cmd *cobra.Command, args []string) {
		date := FlagDate

		for date.Weekday() != time.Monday {
			date = date.AddDate(0, 0, -1)
		}

		endOfWeek := date.AddDate(0, 0, 6)

		from := dateutil.BeginOfDay(date.Year(), date.Month(), date.Day())
		to := dateutil.EndOfDay(endOfWeek.Year(), endOfWeek.Month(), endOfWeek.Day())

		show(from, to, FlagTag)
	},
}

var showDateCmd = &cobra.Command{
	Use:     "date",
	Aliases: []string{"day"},
	Short:   "Show date",
	Long:    `Long description`,
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		FlagDate, err = dateutil.ParseDate(FlagDateString, time.Now())
		return
	},
	Run: func(cmd *cobra.Command, args []string) {
		date := FlagDate

		from := dateutil.BeginOfDay(date.Year(), date.Month(), date.Day())
		to := dateutil.EndOfDay(date.Year(), date.Month(), date.Day())

		show(from, to, FlagTag)
	},
}

func show(from, to time.Time, tag string) {
	db := database.MustOpen(DatabaseFile)
	defer db.Close()

	stmt := db.Joins("JOIN intervals ON trackings.interval_id = intervals.id")

	if !from.IsZero() {
		stmt = stmt.Where("intervals.`from` >= ?", from)
	}

	if !to.IsZero() {
		stmt = stmt.Where("intervals.`to` != ? AND intervals.`to` <= ?", time.Time{}, to)
	}

	if tag != "" {
		stmt = stmt.Where("tag = ?", tag)
	}

	var trackings []tracking.Tracking
	if err := stmt.Find(&trackings).Error; err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tracking.RenderTable(os.Stdout, trackings...)
}

func init() {
	now := time.Now()

	showYearCmd.Flags().IntVarP(&FlagYear, "year", "", now.Year(), "Year")
	showMonthCmd.Flags().IntVarP(&FlagYear, "year", "", now.Year(), "Year")
	showMonthCmd.Flags().IntVarP(&FlagMonth, "month", "", int(now.Month()), "Month")
	showWeekCmd.Flags().IntVarP(&FlagYear, "year", "", now.Year(), "Year")
	showWeekCmd.Flags().IntVarP(&FlagMonth, "month", "", int(now.Month()), "Month")
	showWeekCmd.Flags().StringVarP(&FlagDateString, "date", "", "", "Date")
	showDateCmd.Flags().StringVarP(&FlagDateString, "date", "", "", "Date")

	showCmd.PersistentFlags().StringVarP(&FlagTag, "tag", "t", "", "Tracking tag to use")

	showCmd.AddCommand(showYearCmd)
	showCmd.AddCommand(showMonthCmd)
	showCmd.AddCommand(showWeekCmd)
	showCmd.AddCommand(showDateCmd)

	rootCmd.AddCommand(showCmd)
}
