// Package cmd /*
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
	"os/exec"
)

// weatherCmd represents the weather command
var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "Displays the current weather to you",
	Long:  `Displays the weather to you. This can be at your location or at one you enter`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("moon") {
			moonShow()
		} else {
			location, _ := cmd.Flags().GetString("location")
			if location == "" {
				location = viper.GetString("stdLocation")
			}
			display(weatherShowIn(location))
		}
	},
}

func weatherShow() string {
	return weatherShowIn(viper.GetString("stdLocation"))
}
func weatherShowIn(location string) string {
	command := "wttr.in/" + location + "?format=3"

	cmd2 := exec.Command("curl", command)
	stdout, err := cmd2.StdoutPipe()
	catchError(err)

	if err := cmd2.Start(); err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(stdout)
	catchError(err)

	if err := cmd2.Wait(); err != nil {
		log.Println("ERROR: location unknown")
	}
	return string(data)
}

func moonShow() {
	cmd2 := exec.Command("curl", "wttr.in/moon")
	out, err := cmd2.Output()
	catchError(err)
	display(string(out))
}

func init() {
	rootCmd.AddCommand(weatherCmd)

	// Here you will define your flags and configuration settings.
	weatherCmd.Flags().StringP("location", "l", "", "Gives you the current weather at a given location")
	weatherCmd.Flags().BoolP("moon", "m", false, "Displays the current moon-phase")

}
