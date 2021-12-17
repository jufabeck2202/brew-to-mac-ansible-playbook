package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
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

// options for the command
var exampleOpts = struct {
	id    int
	c     string
	dir   string
	opt   string
	names gcli.Strings
}{}

func main() {
	app := gcli.NewApp()
	app.Version = "1.0.0"
	app.Desc = "A tool to help you manage your brew and cask installs using ansible"
	// app.SetVerbose(gcli.VerbDebug)

	app.Add(&gcli.Command{
		Name: "parse",
		// allow color tag and {$cmd} will be replace to 'demo'
		Desc: "update your ansible notebook using your current installed brew packages <info>parse</> for {$cmd}",
		// ... allow add subcommands
		Aliases: []string{"p"},
		Config: func(c *gcli.Command) {
			// binding options
			// ...
			c.IntOpt(&exampleOpts.id, "id", "", 2, "the id option")
			c.StrOpt(&exampleOpts.c, "config", "c", "value", "the config option")
			// notice `DIRECTORY` will replace to option value type
			c.StrOpt(&exampleOpts.dir, "dir", "d", "", "the `DIRECTORY` option")
			// 支持设置选项短名称
			c.StrOpt(&exampleOpts.opt, "opt", "o", "", "the option message")
			// 支持绑定自定义变量, 但必须实现 flag.Value 接口
			c.VarOpt(&exampleOpts.names, "names", "n", "the option message")

			// binding arguments
			c.AddArg("arg0", "the first argument, is required", true)
			// ...
		},
		Func: Run,
	})
	app.Run(nil)
}

func Run(cmd *gcli.Command, args []string) error {
	brewDump := getBrewDump()
	ansibleFile := readFile(inputFile)
	updatedText := Parse(brewDump, ansibleFile)
	saveToFile(updatedText)
	return nil
}

func Parse(brewDump string, ansibleFile string) string {
	parseBrewFile(brewDump)
	readMacPlayBook(ansibleFile)
	compareSlices(true)
	return updateAnsible(ansibleFile)

}

func getBrewDump() string {
	cmd := "HOMEBREW_NO_AUTO_UPDATE=1 brew bundle dump -q  --file=-"
	out, err := exec.Command("/bin/bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func updateAnsible(input string) string {

	scanner := bufio.NewScanner(strings.NewReader(input))
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
	return totalText
}

func appendToText(text string, packageSuffix string, slice []lineObject) string {
	for _, object := range slice {
		text += packageSuffix + "- " + object.line + "    # " + object.comment + "\n"
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

func parseBrewFile(brewDump string) {

	currentString := string(brewDump)
	scanner := bufio.NewScanner(strings.NewReader(currentString))
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

func readFile(path string) string {
	currentFile, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	return string(currentFile)
}

func readMacPlayBook(input string) {

	scanner := bufio.NewScanner(strings.NewReader(input))

	textSlice := make([]string, 1)
	for scanner.Scan() {
		textSlice = append(textSlice, scanner.Text())
	}
	var extractingTaps, extractingBrews, extractingCasks, extractingMasses bool

	for i, text := range textSlice {
		if text != "" {
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
			newPackage.comment = "added (" + masName + ") " + time.Now().Format("2006-01-02")
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
			newPackage.comment = "added " + time.Now().Format("2006-01-02")
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
			newPackage.comment = "added " + time.Now().Format("2006-01-02")
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
