package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"log"
)

var colors = []string{"Green", "Red", "Blue", "Magenta"}
var welcomeClr = color.New(color.FgGreen)

func catchError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

func displayWelcome() {
	var clr = viper.GetString("welcomeClr")
	switch clr {
	case colors[0]:
		welcomeClr = color.New(color.FgGreen)
	case colors[1]:
		welcomeClr = color.New(color.FgRed)
	case colors[2]:
		welcomeClr = color.New(color.FgBlue)
	case colors[3]:
		welcomeClr = color.New(color.FgMagenta)
	}
	name := viper.GetString("name")
	display("\n")
	_, _ = welcomeClr.Println("	---Welcome Home " + name + "---")
	display("")
}
