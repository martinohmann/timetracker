package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/martinohmann/timetracker/pkg/dateutil"
	"github.com/martinohmann/timetracker/pkg/version"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:               "timetracker",
		Short:             "Track time intervals using tags",
		Version:           version.Version,
		PersistentPreRunE: preFlight,
	}

	dateString      string
	startDateString string
	endDateString   string
	id              int
	month           int
	tag             string
	year            int

	date      time.Time
	startDate time.Time
	endDate   time.Time

	config string
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&config, "config", "", "config file (default is $HOME/.timetracker.yaml)")
	rootCmd.PersistentFlags().StringVar(&startDateString, "start", "", "start date of the interval")
	rootCmd.PersistentFlags().StringVar(&endDateString, "end", "", "end date of the interval")
	rootCmd.PersistentFlags().StringVarP(&tag, "tag", "t", "", "interval tag to use")
	rootCmd.PersistentFlags().String("database", "~/.timetracker.db", "path to sqlite database")
	rootCmd.PersistentFlags().Bool("debug", false, "enable debug output")
	viper.BindPFlag("database", rootCmd.PersistentFlags().Lookup("database"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if config != "" {
		// Use config file from the flag.
		viper.SetConfigFile(config)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".timetracker" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".timetracker")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error while reading config file %s:\n", viper.ConfigFileUsed())
		fmt.Println(err)
		os.Exit(1)
	}
}

// preFlight converts string flags to proper time.Time values
func preFlight(cmd *cobra.Command, args []string) (err error) {
	if startDate, err = dateutil.ParseDate(startDateString, time.Time{}); err != nil {
		return
	}

	endDate, err = dateutil.ParseDate(endDateString, time.Time{})
	return
}

func initializeYearFlag(cmd *cobra.Command) {
	cmd.Flags().IntVar(&year, "year", time.Now().Year(), "filter year")
}

func initializeDateFlag(cmd *cobra.Command) {
	cmd.Flags().StringVar(&dateString, "date", "", "filter date")
}

func initializeIdFlag(cmd *cobra.Command) {
	cmd.Flags().IntVar(&id, "id", 0, "interval ID")
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
