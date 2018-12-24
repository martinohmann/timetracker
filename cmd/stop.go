package cmd

import (
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop an interval",
	Run:   stop,
}

func init() {
	initializeIdFlag(stopCmd)
	rootCmd.AddCommand(stopCmd)
}

func stop(cmd *cobra.Command, args []string) {
	db := database.MustOpen(viper.GetString("database"))
	defer db.Close()

	var err error
	var i interval.Interval

	if tag != "" {
		db = db.Where("tag = ?", tag)
	}

	if id > 0 {
		err = db.First(&i, id).Error
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

	i.Close()

	db.Save(&i)

	interval.RenderTable(cmd.OutOrStdout(), i)
}
