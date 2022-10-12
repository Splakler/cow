/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

var flags = []string{"all", "name", "ip", "wifi"}

// systemCmd represents the system command
var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "Gives you information on your system",
	Long: `Displays all sorts of information about your system. Current functions include:
		--ip
		--name
		--wifi
		--all`,
	Run: func(cmd *cobra.Command, args []string) {
		systemMain(cmd)
	},
}

func systemMain(cmd *cobra.Command) {
	changed := false
	fmt.Println(cmd.Flags())
	for _, element := range flags {
		if cmd.Flag(element).Changed {
			changed = true
			switch element {
			case "all":
				display(sysShowAll())
				break
			case "ip":
				display(sysShowIp())
			case "name":
				display(sysShowName())
			case "wifi":
				display(sysShowWifi())
			}
		}
	}
	if !changed {
		sysShowAll()
	}
}
func init() {
	rootCmd.AddCommand(systemCmd)

	// Here you will define your flags and configuration settings.
	systemCmd.Flags().BoolP("all", "a", false, "shows all system informations")
	systemCmd.Flags().BoolP("ip", "i", false, "shows the current ip adress")
	systemCmd.Flags().BoolP("name", "n", false, "shows the current logged in user")
	systemCmd.Flags().BoolP("wifi", "w", false, "shows information about the current wifi connection")

}
func sysShowAll() string {
	return sysShowName() + "\n" + sysShowIp() + "\n" + sysShowWifi()
}

func sysShowIp() string {
	var cmd2 = exec.Command("hostname", "-i")
	var out, err = cmd2.Output()
	catchError(err)

	return string(out)
}

func sysShowName() string {
	var cmd2 = exec.Command("hostname")
	var out, err = cmd2.Output()
	catchError(err)

	return string(out)
}

func sysShowWifi() string {
	var cmd2 = exec.Command("nmcli", "connection", "show")
	var cmd3 = exec.Command("nmcli", "general", "status")
	var out, err = cmd2.Output()
	var out3, err3 = cmd3.Output()

	catchError(err3)
	catchError(err)
	return string(out3) + "\n" + string(out)
}

func display(input string) {
	fmt.Println(input)
}
