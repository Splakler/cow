// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A detailed Prompt-Interface for your cow!",
	Long:  `A detailed Prompt-Interface for your cow!`,
	Run: func(cmd *cobra.Command, args []string) {
		var promptSelect promptui.Select
		viper.AutomaticEnv()

		displayWelcome()

		for true {
			promptSelect = promptui.Select{
				Label:        "Select Operation",
				Items:        []string{"System", "Weather", "Ping", "Storage", "Settings", "Clear", "Exit"},
				HideSelected: true,
				Size:         10,
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
				displayWelcome()
			case "Settings":
				prgSettings(promptSelect)
			case "Clear":
				clear()
				displayWelcome()
			case "Exit":
				os.Exit(0)
			}
		}
	},
}

func clear() {
	cmd2 := exec.Command("clear")
	cmd2.Stdout = os.Stdout
	_ = cmd2.Run()
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
	displayWelcome()
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
	displayWelcome()
}

func prgPing() {
	promptInput := promptui.Prompt{
		Label:    "Adresse",
		Validate: validatePing,
	}
	result, err := promptInput.Run()
	catchError(err)
	display(pingShow(result))
	displayWelcome()
}

func prgSettings(promptSelect promptui.Select) {
	promptSelect = promptui.Select{
		Label:        "Select Setting",
		Items:        []string{"- Name", "- Location", "- Color", "- Exit"},
		HideSelected: true,
	}
	_, name, err := promptSelect.Run()
	catchError(err)

	if name == "- Exit" {
		return
	}

	var input string
	if name != "- Color" {
		labelStr := "Change " + name[2:]
		promptInput := promptui.Prompt{
			Label: labelStr,
		}
		input, err = promptInput.Run()
		catchError(err)
	} else {
		fmt.Println("Your current Color is: " + viper.GetString("welcomeClr"))
		promptSelect = promptui.Select{
			Label:        "Select a Color",
			Items:        colors,
			HideSelected: true,
		}
		_, input, err = promptSelect.Run()
		catchError(err)
	}

	switch name {
	case "- Name":
		changeConfig("name", input)
	case "- Location":
		changeConfig("stdLocation", input)
	case "- Color":
		changeConfig("welcomeClr", input)
	case "- Exit":
		break
	}
	displayWelcome()
}

func init() {
	rootCmd.AddCommand(startCmd)

}
