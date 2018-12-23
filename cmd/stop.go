package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/tracking"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop time tracking",
	Long:  `Long description`,
	Run: func(cmd *cobra.Command, args []string) {
		db := database.MustOpen(FlagDatabase)
		defer db.Close()

		var err error
		var t tracking.Tracking

		if FlagTag != "" {
			db = db.Where("tag = ?", FlagTag)
		}

		if FlagID > 0 {
			err = db.First(&t, FlagID).Error
		} else {
			err = db.
				Joins("JOIN intervals ON trackings.interval_id = intervals.id").
				Where("intervals.end = ?", time.Time{}).
				Last(&t).Error
		}

		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("nothing to stop")
			os.Exit(0)
		} else if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if t.Finished {
			fmt.Printf("tracking %d is already finished\n", t.ID)
			os.Exit(1)
		}

		t.Interval.End = time.Now()
		t.Finished = true

		db.Save(&t)

		tracking.RenderTable(os.Stdout, t)
	},
}

func init() {
	stopCmd.Flags().StringVarP(&FlagTag, "tag", "t", "", "Tracking tag to use")
	stopCmd.Flags().IntVarP(&FlagID, "id", "", 0, "Tracking ID")
	rootCmd.AddCommand(stopCmd)
}
