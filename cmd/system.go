/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var flags = []string{"all", "name", "ip", "wifi"}

// systemCmd represents the system command
var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		changed := false
		for _, element := range flags {
			if cmd.Flag(element).Changed {
				changed = true
				switch element {
				case "all":
					showAll()
					break
				case "ip":
					showIp()
				case "name":
					showName()
				case "wifi":
					showWifi()
				}
			}
		}
		if !changed {
			showAll()
		}
	},
}

func init() {
	rootCmd.AddCommand(systemCmd)

	// Here you will define your flags and configuration settings.
	systemCmd.Flags().BoolP("all", "a", false, "shows all system informations")
	systemCmd.Flags().BoolP("ip", "i", false, "shows the current ip adress")
	systemCmd.Flags().BoolP("name", "n", false, "shows the current logged in user")
	systemCmd.Flags().BoolP("wifi", "w", false, "shows information about the current wifi connection")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// systemCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// systemCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func showAll() {
	showName()
	fmt.Println()
	showIp()
	fmt.Println()
	showWifi()
}
func showIp() {
	var cmd2 = exec.Command("hostname", "-i")
	var out, err = cmd2.Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", out)
}

func showName() {
	var cmd2 = exec.Command("hostname")
	var out, err = cmd2.Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", out)
}
func showWifi() {
	var cmd2 = exec.Command("nmcli", "connection", "show")
	var cmd3 = exec.Command("nmcli", "general", "status")
	var out, err = cmd2.Output()
	var out3, err3 = cmd3.Output()

	if err3 != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out3)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
}
