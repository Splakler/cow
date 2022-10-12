// Package cmd /*
package cmd

import (
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A detailed Prompt-Interface for your cow!",
	Long:  `A detailed Prompt-Interface for your cow!`,
	Run: func(cmd *cobra.Command, args []string) {
		var promptSelect promptui.Select
		viper.AutomaticEnv()
		clWelcome := color.New(color.FgGreen)
		name := viper.GetString("USER")
		display("\n")
		_, _ = clWelcome.Println("	---Welcome Home " + name + "---")
		for true {
			promptSelect = promptui.Select{
				Label:        "Select Operation",
				Items:        []string{"System", "Weather", "Ping", "Storage", "Exit"},
				HideSelected: true,
			}
			_, result, err := promptSelect.Run()

			catchError(err)

			switch result {
			case "System":
				prgSystem(promptSelect)
			case "Weather":
				prgWeather(promptSelect)
			case "Ping":
				prgPing()
			case "Storage":
				showStorage()
			case "Exit":
				os.Exit(0)
			}
		}
	},
}

func prgSystem(promptSelect promptui.Select) {
	promptSelect = promptui.Select{
		Label:        "Select specification",
		Items:        []string{"- All", "- Ip", "- Name", "- Wifi", "- Exit"},
		HideSelected: true,
	}
	_, result, err := promptSelect.Run()
	catchError(err)

	switch result {
	case "- All":
		display(sysShowAll())
	case "- Ip":
		display(sysShowIp())
	case "- Name":
		display(sysShowName())
	case "- Wifi":
		display(sysShowWifi())
	case "- Exit":
		break
	}
}

func prgWeather(promptSelect promptui.Select) {
	display(weatherShow())
	for true {
		promptSelect = promptui.Select{
			Items:        []string{"- Change Location", "- Show Moon-phase", "- Exit"},
			HideSelected: true,
			HideHelp:     true,
			Label:        "Weather",
		}
		_, result, err := promptSelect.Run()
		catchError(err)

		switch result {
		case "- Change Location":
			promptInput := promptui.Prompt{Label: "Location"}
			result, err = promptInput.Run()
			catchError(err)
			display(weatherShowIn(result))
		case "- Show Moon-phase":
			moonShow()

		case "- Exit":
			break
		}
		if result == "- Exit" {
			break
		}
	}
}

func prgPing() {
	promptInput := promptui.Prompt{
		Label:    "Adresse",
		Validate: validatePing,
	}
	result, err := promptInput.Run()
	catchError(err)
	display(pingShow(result))
}

func init() {
	rootCmd.AddCommand(startCmd)

}

func catchError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}
