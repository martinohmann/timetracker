package cmd

import (
	"fmt"
	"os"

	"github.com/martinohmann/timetracker/pkg/database"
	"github.com/martinohmann/timetracker/pkg/interval"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "Delete an interval",
	Run:     del,
}

func init() {
	initializeIdFlag(deleteCmd)
	deleteCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(deleteCmd)
}

func del(cmd *cobra.Command, args []string) {
	db := database.MustOpen(viper.GetString("database"))
	defer db.Close()

	var i interval.Interval

	if err := db.First(&i, id).Error; err != nil {
		cmd.Println(err)
		os.Exit(1)
	}

	db.Delete(&i)

	interval.RenderTable(cmd.OutOrStdout(), i)

	fmt.Println("deleted")
}
