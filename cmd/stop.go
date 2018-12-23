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
		db := database.MustOpen(DatabaseFile)
		defer db.Close()

		var err error
		var t tracking.Tracking

		if FlagID > 0 {
			err = db.First(&t, FlagID).Error
		} else {
			var tm time.Time
			err = db.
				Joins("JOIN intervals ON trackings.interval_id = intervals.id").
				Where("intervals.`to` = ?", tm).
				Last(&t).Error
		}

		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("nothing to stop")
			os.Exit(0)
		} else if err != nil {
			panic(err)
		}

		if t.Finished {
			fmt.Printf("tracking %d is already finished\n", t.ID)
			os.Exit(1)
		}

		t.Interval.To = time.Now()
		t.Finished = true

		db.Save(&t)

		tracking.RenderTable(os.Stdout, t)
	},
}

func init() {
	stopCmd.Flags().IntVarP(&FlagID, "id", "", 0, "Tracking ID")
	rootCmd.AddCommand(stopCmd)
}
