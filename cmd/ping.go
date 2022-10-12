/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os/exec"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "lets you ping stuff",
	Long:  `Works just like normal ping command. Enter a URL or IP-Adress and let it do the rest`,
	Run: func(cmd *cobra.Command, args []string) {
		display(pingShow(args[0]))
	},
}

func pingShow(input string) string {
	var cmd2 = exec.Command("ping", input, "-c 3")
	var out, err = cmd2.Output()
	catchError(err)

	return string(out)
}

func validatePing(input string) error {
	var cmd2 = exec.Command("ping", input, "-c 1", "-W 0.5", "-s 2")
	var _, err = cmd2.Output()

	return err
}

func init() {
	rootCmd.AddCommand(pingCmd)

}
