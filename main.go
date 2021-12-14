package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/gookit/gcli/v3"
	"github.com/gookit/gcli/v3/interact"
)

var tapSliceBrew = make([][]string, 0)
var brewSliceBrew = make([][]string, 0)
var caskSliceBrew = make([][]string, 0)
var masSliceBrew = make([][]string, 0)

var tapSlice = make([]lineObject, 0)
var brewSlice = make([]lineObject, 0)
var caskSlice = make([]lineObject, 0)
var masSlice = make([]lineObject, 0)

var brewNotInstalled = make([]lineObject, 0)
var caskNotInstalled = make([]lineObject, 0)
var tabsNotInstalled = make([]lineObject, 0)
var masNotInstalled = make([]lineObject, 0)

var inputFile = "./default.config.empty.yaml"
var outputFile = "./default.config2.yaml"

type lineObject struct {
	line       string
	lineNumber int
	comment    string
}

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
	readBrewFile()
	readMacPlayBook()
	compareSlices(true)
	appendToFile()

}
func appendToFile() {

	currentFile, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	currentString := string(currentFile)
	scanner := bufio.NewScanner(strings.NewReader(currentString))
	var extractingTaps, extractingBrews, extractingCasks, extractingMasses bool
	totalText := ""
	packageSuffix := ""
	for scanner.Scan() {
		line := scanner.Text()
		strippedLine := strings.ReplaceAll(line, " ", "")
		//skip empty line
		if strippedLine == "" {
			totalText += line + "\n"
			continue
		}
		//skip comments
		if strippedLine[0:1] == "#" {
			totalText += line + "\n"
			continue
		}
		//skip current package
		if strippedLine[0:1] == "-" {
			packageSuffix = leadingWhitespace(line)
			totalText += line + "\n"
			continue
		}
		if extractingBrews {
			totalText = appendToText(totalText, packageSuffix, brewNotInstalled)
			extractingBrews = false
		}
		if extractingCasks {
			totalText = appendToText(totalText, packageSuffix, caskNotInstalled)
			extractingCasks = false
		}
		if extractingTaps {
			totalText = appendToText(totalText, packageSuffix, tabsNotInstalled)
			extractingTaps = false
		}
		if extractingMasses {
			totalText = appendToText(totalText, packageSuffix, masNotInstalled)
			extractingMasses = false
		}
		totalText += line + "\n"
		if strings.Contains(strippedLine, strings.TrimSpace("homebrew_installed_packages:")) {
			extractingTaps, extractingBrews, extractingCasks, extractingMasses = false, true, false, false
			continue
		}
		if strings.Contains(strippedLine, strings.TrimSpace("homebrew_taps:")) {
			extractingTaps, extractingBrews, extractingCasks, extractingMasses = true, false, false, false
			continue
		}
		if strings.Contains(strippedLine, strings.TrimSpace("homebrew_cask_apps:")) {
			extractingTaps, extractingBrews, extractingCasks, extractingMasses = false, false, true, false
			continue
		}
		if strings.Contains(strippedLine, strings.TrimSpace("mas_installed_apps:")) {
			extractingTaps, extractingBrews, extractingCasks, extractingMasses = false, false, false, true
			continue
		}

	}
	saveToFile(totalText)

}
func appendToText(text string, packageSuffix string, slice []lineObject) string {
	for _, object := range slice {
		text += packageSuffix + "- " + object.line + "    #" + object.comment + "\n"
	}
	return text
}
func saveToFile(text string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	} else {
		file.WriteString(text)
	}
	file.Close()
	return nil
}

func leadingWhitespace(line string) string {
	leading := ""
	for i := 0; i < len(line); i++ {
		if line[i] != ' ' {
			return leading
		}
		leading += " "
	}
	return leading
}
func readBrewFile() {

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
}
func readMacPlayBook() {

	file, err := os.Open(inputFile)

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

	for i, text := range textSlice {
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
				brewSlice = trimAndAddToSlice(line, i, brewSlice)
			} else if extractingTaps {
				tapSlice = trimAndAddToSlice(line, i, tapSlice)

			} else if extractingCasks {
				caskSlice = trimAndAddToSlice(line, i, caskSlice)

			} else if extractingMasses {
				masSlice = trimAndAddToSlice(line, i, masSlice)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
func trimAndAddToSlice(line string, number int, slice []lineObject) []lineObject {
	if line[0:1] != "-" {
		return slice
	}
	return append(slice, lineObject{line[1:], number, ""})

}

func getQuotedStrings1(s string) []string {
	var re1 = regexp.MustCompile(`"(.*?)"`)
	ms := re1.FindAllStringSubmatch(s, -1)
	ss := make([]string, len(ms))
	for i, m := range ms {
		ss[i] = m[1]
	}
	return ss
}

func compareSlices(addAll bool) {

	color.Info.Println("Mas Slices:")
	for _, line := range masSliceBrew {
		totalString := strings.Join(line[:], " ")
		masName := getQuotedStrings1(totalString)[0]
		packageId := line[len(line)-1]
		alreadyInstalled, newPackage := stringInSlice(packageId, masSlice)
		if !alreadyInstalled {
			newPackage.comment = packageId

			if addAll {
				masNotInstalled = append(masNotInstalled, newPackage)
				continue
			}

			install := interact.Confirm(color.Cyan.Text("\t - Add " + masName + " to Ansible (" + packageId + ") ?"))
			if install {
				masNotInstalled = append(masNotInstalled, newPackage)
			}
		}

	}
	color.Info.Println("Brew Packages: ")
	for _, line := range brewSliceBrew {
		cleanLine := strings.ReplaceAll(strings.ReplaceAll(line[1], "\"", ""), ",", "")
		alreadyInstalled, newPackage := stringInSlice(cleanLine, brewSlice)
		if !alreadyInstalled {
			newPackage.comment = "added " + time.Now().Format("2006-01-02")
			if addAll {
				brewNotInstalled = append(brewNotInstalled, newPackage)
				continue
			}
			install := interact.Confirm(color.Cyan.Text("\t - Add " + cleanLine + " to Ansible?"))
			if install {
				brewNotInstalled = append(brewNotInstalled, newPackage)
			}
		}

	}

	color.Info.Println("Brew Cask Packages: ")
	for _, line := range caskSliceBrew {
		cleanLine := strings.ReplaceAll(strings.ReplaceAll(line[1], "\"", ""), ",", "")
		alreadyInstalled, newPackage := stringInSlice(cleanLine, caskSlice)
		if !alreadyInstalled {
			if addAll {
				caskNotInstalled = append(caskNotInstalled, newPackage)
				continue
			}
			install := interact.Confirm(color.Cyan.Text("\t - Add " + cleanLine + " to Ansible?"))
			if install {
				caskNotInstalled = append(caskNotInstalled, newPackage)
			}
		}

	}

	color.Info.Println("Brew Taps: ")
	for _, line := range tapSliceBrew {
		cleanLine := strings.ReplaceAll(strings.ReplaceAll(line[1], "\"", ""), ",", "")
		alreadyInstalled, newPackage := stringInSlice(cleanLine, tapSlice)
		if !alreadyInstalled {
			if addAll {
				tabsNotInstalled = append(tabsNotInstalled, newPackage)
				continue
			}
			install := interact.Confirm(color.Cyan.Text("\t - Add " + cleanLine + " to Ansible?"))
			if install {
				tabsNotInstalled = append(tabsNotInstalled, newPackage)
			}
		}

	}
	color.Info.Println("New Brew Packages: " + strconv.Itoa(len(brewNotInstalled)))
	color.Info.Println("New Brew Cask Packages: " + strconv.Itoa(len(caskNotInstalled)))
	color.Info.Println("New Brew Taps: " + strconv.Itoa(len(tabsNotInstalled)))
	color.Info.Println("New Mas Packages: " + strconv.Itoa(len(masNotInstalled)))
}

func stringInSlice(a string, lineObjects []lineObject) (bool, lineObject) {
	for _, b := range lineObjects {
		if b.line == a {
			return true, b

		}
	}
	return false, lineObject{a, 0, ""}

}

func combineAndSort(slices ...[]lineObject) []lineObject {
	combined := make([]lineObject, 0)
	for _, slice := range slices {
		combined = append(combined, slice...)
	}
	sort.Slice(combined, func(i, j int) bool {
		return combined[i].line < combined[j].line

	})
	return combined
}
