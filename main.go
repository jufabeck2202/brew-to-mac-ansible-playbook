package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gookit/gcli/v3"
)

var tapSliceBrew = make([][]string, 1)
var brewSliceBrew = make([][]string, 1)
var caskSliceBrew = make([][]string, 1)
var masSliceBrew = make([][]string, 1)
var tapSlice = make([]string, 1)
var brewSlice = make([]string, 1)
var caskSlice = make([]string, 1)
var masSlice = make([]string, 1)

func main() {
	app := gcli.NewApp()
	app.Version = "1.0.3"
	app.Desc = "this is my cli application"
	// app.SetVerbose(gcli.VerbDebug)

	app.Add(&gcli.Command{
		Name: "demo",
		// allow color tag and {$cmd} will be replace to 'demo'
		Desc: "this is a description <info>message</> for {$cmd}",
		Subs: []*gcli.Command{
			// ... allow add subcommands
		},
		Aliases: []string{"dm"},
		Func: func(cmd *gcli.Command, args []string) error {
			gcli.Print("hello, in the demo command\n")
			return nil
		},
	})
	// app.Run(nil)

	file, err := os.Open("./Brewfile")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	for scanner.Scan() {
		lineArray := strings.Fields(scanner.Text())

		switch lineArray[0] {
		case "tap":
			tapSliceBrew = append(tapSliceBrew, lineArray)
		case "brew":
			brewSliceBrew = append(brewSliceBrew, lineArray)
		case "cask":
			caskSliceBrew = append(caskSliceBrew, lineArray)
		case "mas":
			masSliceBrew = append(masSliceBrew, lineArray)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	readMacPlayBook()
}

func readMacPlayBook() {

	file, err := os.Open("./default.config.yaml")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	textSlice := make([]string, 1)
	for scanner.Scan() {
		textSlice = append(textSlice, scanner.Text())
	}
	var extractingTaps, extractingBrews, extractingCasks, extractingMasses bool

	for _, text := range textSlice {
		if text != "" {
			// fmt.Print(text)
			// fmt.Println(strings.ReplaceAll(text, " ", "")[0:1] != "-")
			line := strings.ReplaceAll(text, " ", "")
			if line[0:1] != "-" && line[0:1] != "#" {
				extractingBrews = false
				extractingTaps = false
				extractingCasks = false
				extractingMasses = false
			}
			if strings.Contains(line, strings.TrimSpace("homebrew_installed_packages")) {
				extractingBrews = true
			}
			if strings.Contains(line, strings.TrimSpace("homebrew_taps")) {
				extractingTaps = true
			}
			if strings.Contains(line, strings.TrimSpace("homebrew_cask_apps")) {
				extractingCasks = true
			}
			if strings.Contains(line, strings.TrimSpace("mas_installed_apps")) {
				extractingMasses = true
			}

			if extractingBrews {
				brewSlice = trimAndAddToSlice(line, brewSlice)
			} else if extractingTaps {
				tapSlice = trimAndAddToSlice(line, tapSlice)
			} else if extractingCasks {
				caskSlice = trimAndAddToSlice(line, caskSlice)
			} else if extractingMasses {
				masSlice = trimAndAddToSlice(line, masSlice)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(masSlice)

}
func trimAndAddToSlice(line string, slice []string) []string {
	if line[0:1] != "-" {
		return slice
	}
	return append(slice, line[1:])

}
