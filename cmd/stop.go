package cmd

import (
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop an interval",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.MustOpen(FlagDatabase, FlagDebug)
		defer db.Close()

		var err error
		var i interval.Interval

		if FlagTag != "" {
			db = db.Where("tag = ?", FlagTag)
		}

		if FlagID > 0 {
			err = db.First(&i, FlagID).Error
		} else {
			err = db.Where("end = ?", time.Time{}).
				Last(&i).Error
		}

		if gorm.IsRecordNotFoundError(err) {
			cmd.Println("nothing to stop")
			os.Exit(0)
		} else if err != nil {
			cmd.Println(err)
			os.Exit(1)
		}

		if i.IsClosed() {
			cmd.Printf("interval %d is already closed\n", i.ID)
			os.Exit(1)
		}

		i.End = time.Now()

		db.Save(&i)

		interval.RenderTable(cmd.OutOrStdout(), i)
	},
}

func init() {
	stopCmd.Flags().IntVarP(&FlagID, "id", "", 0, "interval ID")
	rootCmd.AddCommand(stopCmd)
}
