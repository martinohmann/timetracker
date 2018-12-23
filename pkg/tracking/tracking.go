package tracking

import (
	"fmt"
	"io"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/olekukonko/tablewriter"
)

const DefaultTag = "default"

type Tracking struct {
	gorm.Model
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

	var totalDuration time.Duration

	for _, t := range trackings {
		from := t.Interval.From.Format("2006-01-02 15:04:05")

		var to, duration string

		if t.Interval.To.IsZero() {
			to = "-"
		} else {
			to = t.Interval.To.Format("2006-01-02 15:04:05")
		}

		duration = t.Interval.Duration().String()
		totalDuration += t.Interval.Duration()

		table.Append([]string{
			fmt.Sprintf("%d", t.ID),
			t.Tag,
			from,
			to,
			duration,
		})
	}

	table.SetAutoFormatHeaders(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetFooterAlignment(tablewriter.ALIGN_RIGHT)

	if len(trackings) > 1 {
		table.SetFooter([]string{"", "", "", "Total", totalDuration.String()})
	}

	table.Render()
}
