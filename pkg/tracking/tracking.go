package tracking

import (
	"io"
	"sort"
	"strconv"
	"time"

	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/olekukonko/tablewriter"
)

const (
	DefaultTag = "default"

	TimeFormat = "2006/01/02 15:04:05"
)

type Tracking struct {
	ID         int
	Tag        string
	Interval   interval.Interval
	IntervalID int
	Finished   bool
}

func RenderTable(w io.Writer, trackings ...Tracking) {
	if len(trackings) == 0 {
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

	sort.Sort(Collection(trackings))

	for _, t := range trackings {
		start := t.Interval.Start.Format(TimeFormat)
		end := "-"
		if !t.Interval.End.IsZero() {
			end = t.Interval.End.Format(TimeFormat)
		}

		duration := t.Interval.Duration().Truncate(time.Second)
		totalDuration += t.Interval.Duration().Truncate(time.Second)

		table.Append([]string{
			strconv.Itoa(int(t.ID)),
			t.Tag,
			start,
			end,
			duration.String(),
		})
	}

	table.SetAutoFormatHeaders(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetFooterAlignment(tablewriter.ALIGN_RIGHT)

	if len(trackings) > 1 {
		totalDuration = totalDuration.Truncate(time.Second)
		table.SetFooter([]string{"", "", "", "Total", totalDuration.String()})
	}

	table.Render()
}
