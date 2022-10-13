package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
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

func display(input string) {
	fmt.Println(input)
}

func writeTable(list, names []string) {
	table := tablewriter.NewWriter(os.Stdout)
	if names != nil {
		table.SetHeader(names)
	}

	for _, listElem := range list {
		row := strings.Split(listElem, "\n")
		for _, rowElem := range row {
			rowElem = strings.TrimSpace(rowElem)
			for strings.Contains(rowElem, "   ") {
				rowElem = strings.Replace(rowElem, "   ", "  ", -1)
			}
			table.Append(strings.Split(rowElem, "  "))
		}
	}
	table.Render()
}
