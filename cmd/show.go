package cmd

import (
	"time"

	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/dateutil"
	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/martinohmann/timetracker/pkg/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	showCmd = &cobra.Command{
		Use:               "show [tag]",
		Aliases:           []string{"s", "status"},
		Short:             "Show all intervals",
		PersistentPreRunE: preRunE(parseTagArg, parseDateRange),
		Run:               show,
	}

	showYearCmd = &cobra.Command{
		Use:     "year [tag]",
		Aliases: []string{"y"},
		Short:   "Show year's intervals",
		PreRun:  showYear,
		Run:     show,
	}

	showMonthCmd = &cobra.Command{
		Use:     "month [tag]",
		Aliases: []string{"m"},
		Short:   "Show month's intervals",
		PreRun:  showMonth,
		Run:     show,
	}

	showWeekCmd = &cobra.Command{
		Use:     "week [tag]",
		Aliases: []string{"w"},
		Short:   "Show week's intervals",
		PreRunE: showWeek,
		Run:     show,
	}

	showDateCmd = &cobra.Command{
		Use:     "date [tag]",
		Aliases: []string{"d", "day"},
		Short:   "Show date's intervals",
		PreRunE: showDate,
		Run:     show,
	}
)

func init() {
	initializeDateRangeFlags(showCmd)
	initializeTagFlag(showCmd)

	showMonthCmd.Flags().IntVar(&month, "month", int(time.Now().Month()), "filter month")

	initializeYearFlag(showYearCmd)
	initializeYearFlag(showMonthCmd)
	initializeDateFlag(showWeekCmd)
	initializeDateFlag(showDateCmd)

	showCmd.AddCommand(showYearCmd)
	showCmd.AddCommand(showMonthCmd)
	showCmd.AddCommand(showWeekCmd)
	showCmd.AddCommand(showDateCmd)

	rootCmd.AddCommand(showCmd)
}

func showYear(cmd *cobra.Command, args []string) {
	startDate = dateutil.BeginOfDay(year, time.January, 1)
	endDate = startDate.AddDate(1, 0, 0)
}

func showMonth(cmd *cobra.Command, args []string) {
	startDate = dateutil.BeginOfDay(year, time.Month(month), 1)
	endDate = startDate.AddDate(0, 1, 0)
}

func showWeek(cmd *cobra.Command, args []string) (err error) {
	if date, err = dateutil.ParseDate(dateString, time.Now()); err != nil {
		return
	}

	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, -1)
	}

	startDate = dateutil.BeginOfDay(date.Year(), date.Month(), date.Day())
	endDate = startDate.AddDate(0, 0, 7)

	return
}

func showDate(cmd *cobra.Command, args []string) (err error) {
	if date, err = dateutil.ParseDate(dateString, time.Now()); err != nil {
		return
	}

	startDate = dateutil.BeginOfDay(date.Year(), date.Month(), date.Day())
	endDate = startDate.AddDate(0, 0, 1)

	return
}

func show(cmd *cobra.Command, args []string) {
	database.Init(viper.GetString("database"))
	defer database.Close()

	intervals, err := database.FindIntervalsByCriteria(interval.Interval{
		Tag:   tag,
		Start: startDate,
		End:   endDate,
	})
	exitOnError(err)

	table.Render(cmd.OutOrStdout(), intervals...)
}
