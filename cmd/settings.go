/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// settingsCmd represents the settings command
var settingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var flags = []string{"name", "location", "color"}
		for _, element := range flags {
			if cmd.Flag(element).Changed {
				input, _ := cmd.Flags().GetString(element)
				switch element {
				case "name":
					changeConfig("name", input)
					break
				case "location":
					changeConfig("stdLocation", input)
				case "color":
					changeConfig("welcomeClr", input)
				}
			}
		}
	},
}

func changeConfig(name, value string) {
	viper.Set(name, value)
	fmt.Println("Set " + name + " to " + value)
	_ = viper.WriteConfig()
}

func init() {
	rootCmd.AddCommand(settingsCmd)
	settingsCmd.Flags().StringP("name", "n", viper.GetString("name"), "Changes your name in the programm")
	settingsCmd.Flags().StringP("location", "l", viper.GetString("stdLocation"), "Changes your standart location")
	settingsCmd.Flags().StringP("color", "c", viper.GetString("welcomeClr"), "Changes the color of your welcome message")
}
