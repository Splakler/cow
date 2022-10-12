/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os/exec"
)

// weatherCmd represents the weather command
var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		location, _ := cmd.Flags().GetString("location")
		command := "wttr.in/" + location + "?format=3"

		cmd2 := exec.Command("curl", command)
		stdout, err := cmd2.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}

		if err := cmd2.Start(); err != nil {
			log.Fatal(err)
		}

		data, err := ioutil.ReadAll(stdout)

		if err != nil {
			log.Fatal(err)
		}

		if err := cmd2.Wait(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", string(data))
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)

	// Here you will define your flags and configuration settings.
	weatherCmd.Flags().StringP("location", "l", "", "Gives you the current weather at a given location")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// weatherCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// weatherCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
