package interval

import (
	"io"
	"sort"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

const (
	TimeFormat = "2006/01/02 15:04:05"
)

func RenderTable(w io.Writer, intervals ...Interval) {
	if len(intervals) == 0 {
		return
	}

	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"ID", "Tag", "Start", "End", "Duration"})
	table.SetColumnAlignment([]int{
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_RIGHT,
	})

	var totalDuration time.Duration

	sort.Sort(SortByDate(intervals))

	for _, i := range intervals {
		start := i.Start.Format(TimeFormat)
		end := "-"
		if i.IsClosed() {
			end = i.End.Format(TimeFormat)
		}

		tag := i.Tag
		if tag == "" {
			tag = "-"
		}
		duration := i.Duration().Truncate(time.Second)
		totalDuration += i.Duration().Truncate(time.Second)

		table.Append([]string{
			strconv.Itoa(int(i.ID)),
			tag,
			start,
			end,
			duration.String(),
		})
	}

	table.SetAutoFormatHeaders(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetFooterAlignment(tablewriter.ALIGN_RIGHT)

	if len(intervals) > 1 {
		totalDuration = totalDuration.Truncate(time.Second)
		table.SetFooter([]string{"", "", "", "Total", totalDuration.String()})
	}

	table.Render()
}
