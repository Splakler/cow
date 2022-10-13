// Package cmd /*
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cow",
	Short: "C.ustom O.perations W.indow",
	Long: `
 .----------------.  .----------------.  .----------------. 
| .--------------. || .--------------. || .--------------. |
| |     ______   | || |     ____     | || | _____  _____ | |
| |   .' ___  |  | || |   .'    '.   | || ||_   _||_   _|| |
| |  / .'   \_|  | || |  /  .--.  \  | || |  | | /\ | |  | |
| |  | |         | || |  | |    | |  | || |  | |/  \| |  | |
| |  \ '.___.'\  | || |  \  '--'  /  | || |  |   /\   |  | |
| |   '._____.'  | || |   '.____.'   | || |  |__/  \__|  | |
| |              | || |              | || |              | |
| '--------------' || '--------------' || '--------------' |
 '----------------'  '----------------'  '----------------'
       Custom            Operations            Window

is used to give you all sorts of informations about your system.
for example your IP-Adress, current logged in user, remaining storage, Internet connection
and so much more!`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigFile("config.yaml")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
	}
}
