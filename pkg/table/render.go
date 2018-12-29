package table

import (
	"io"
	"sort"
	"strconv"
	"time"

	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/olekukonko/tablewriter"
)

// TimeFormat defines the time format for the table
var TimeFormat = "2006/01/02 15:04:05"

// Render renders a table of intervals and writes the result to w
func Render(w io.Writer, intervals ...interval.Interval) {
	if len(intervals) == 0 {
		return
	}

	sort.Sort(interval.SortByDate(intervals))

	table := tablewriter.NewWriter(w)
	configureTable(table)

	table.SetHeader([]string{"ID", "Tag", "Start", "End", "Duration"})

	var start, end, tag, humanDuration string
	var duration, total time.Duration

	for _, i := range intervals {
		start = i.Start.Format(TimeFormat)
		if i.IsClosed() {
			end = i.End.Format(TimeFormat)
		} else {
			end = "open"
		}

		if i.Tag != "" {
			tag = i.Tag
		} else {
			tag = "<empty>"
		}

		duration = i.Duration()

		if duration > 0 {
			total += duration
			humanDuration = formatDuration(duration)
		} else {
			humanDuration = "not started"
		}

		table.Append([]string{
			strconv.Itoa(int(i.ID)),
			tag,
			start,
			end,
			humanDuration,
		})
	}

	if len(intervals) > 1 {
		table.SetFooter([]string{
			"",
			"",
			"",
			"Total",
			formatDuration(total),
		})
	}

	table.Render()
}

func formatDuration(d time.Duration) string {
	return d.Truncate(time.Second).String()
}

// configureTable set table formatting options
func configureTable(table *tablewriter.Table) {
	table.SetColumnAlignment([]int{
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_RIGHT,
	})
	table.SetAutoFormatHeaders(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetFooterAlignment(tablewriter.ALIGN_RIGHT)
}
