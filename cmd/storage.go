// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os/exec"
	"strconv"
	"strings"
)

// storageCmd represents the storage command
var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Gives you information about your storage usage",
	Long:  `Shows you a graphical representation of your Storage in GigaBytes`,
	Run: func(cmd *cobra.Command, args []string) {
		showStorage()
	},
}

func showStorage() {
	display("")
	writeTable([]string{convertStorage(strings.Split(getStorage(), "\n")[1])}, []string{"Filesystem", "Size", "Used", "Avail", "Used %", "Mounted on"})
	display("")

	storage := calcStorage()
	full, total := storage[1], storage[0]
	length := 40
	used := int((float32(full) / (float32(total))) * (float32(length)))
	var usedStr, unUsedStr string
	for i := 1; i <= used; i++ {
		usedStr = usedStr + "-"
	}
	for k := used; k <= length; k++ {
		unUsedStr = unUsedStr + " "
	}
	cUsed := color.New(color.BgRed)
	cUnUsed := color.New(color.BgGreen)
	cPercent := color.New(color.FgRed)
	_, _ = cUsed.Print(usedStr)
	_, _ = cUnUsed.Print(unUsedStr)
	_, _ = cPercent.Println("	" + strconv.Itoa(storage[3]) + "% used")
	display("")
}

// Size  Used Avail Use%
func calcStorage() []int {
	out := getStorage()

	lines := strings.Split(out, "\n")
	outStr := lines[1]
	lines = strings.Fields(outStr)
	outStr = fmt.Sprintf("%s/%s/%s/%s", lines[1], lines[2], lines[3], lines[4])
	outStr = strings.Replace(outStr, "G", "", -1)
	outStr = strings.Replace(outStr, "%", "", -1)
	lines = strings.Split(outStr, "/")

	var sizeStrg, usedStrg, availStrg, prctStrg int
	sizeStrg, _ = strconv.Atoi(lines[0])
	usedStrg, _ = strconv.Atoi(lines[1])
	availStrg, _ = strconv.Atoi(lines[2])
	prctStrg, _ = strconv.Atoi(lines[3])

	return []int{sizeStrg, usedStrg, availStrg, prctStrg}
}

func getStorage() string {
	cmd2 := exec.Command("df", "/dev/sda3", "-h")
	out, err := cmd2.Output()
	catchError(err)
	return string(out)
}
func init() {
	rootCmd.AddCommand(storageCmd)
}

func convertStorage(input string) string {
	for strings.Contains(input, "  ") {
		input = strings.Replace(input, "  ", " ", -1)
	}
	return strings.Replace(input, " ", "  ", -1)
}
