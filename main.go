package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const considerLineStartingWith = "###"
const ignoreTag = "@"

var re = regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2})-(\w{3}) :(\d{4})-(\d{4})(-(\d+)m)?$`)

func main() {
	args := os.Args[1:]
	if args == nil || len(args) == 0 || len(args) > 1 {
		fmt.Println("syntax: mdtimesheet ~/filename.md - missing argument")
		os.Exit(0)
	}
	path := strings.Trim(args[0], " ")
	if path == "" {
		fmt.Println("syntax: mdtimesheet ~/filename.md - invalid filepath")
		os.Exit(0)
	}
	contents := readFile(path)
	lines := strings.Split(contents, "\n")
	totalMinutes := 0
	for _, line := range lines {
		if len(line) > 3 && strings.HasPrefix(line, considerLineStartingWith) {
			if strings.HasPrefix(line[4:], ignoreTag) {
				continue
			}
			totalMinutes += parseLine(line[4:])
		}
	}
	color.Set(color.FgCyan)
	fmt.Println("-----------------------------------------")
	fmt.Println("Total minutes spent", totalMinutes, "minutes")
	dur, err := time.ParseDuration(strconv.Itoa(totalMinutes) + "m")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("Total duration spent", dur.String())
	fmt.Println("")
	color.Unset()
}

func readFile(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(bytes)
}

func parseLine(l string) int {
	fmt.Println("\n" + l)
	res := re.FindAllStringSubmatch(l, -1)
	if len(res) == 0 {
		color.Set(color.FgHiRed)
		fmt.Println("\tWARN problem parsing line:", l)
		color.Unset()
		return 0
	}
	if res == nil || res[0] == nil {
		color.Set(color.FgHiRed)
		fmt.Println("\tWARN Bad Line", res)
		color.Unset()
		return 0
	}
	info := res[0]
	if len(info) == 0 {
		color.Set(color.FgHiRed)
		fmt.Println("\tWARN Bad Line Info", res)
		color.Unset()
		return 0
	}

	t1, terr1 := time.Parse("2006-01-02 15:04", info[1]+"-"+info[2]+"-"+info[3]+" "+info[5][:2]+":"+info[5][2:4])
	t2, terr2 := time.Parse("2006-01-02 15:04", info[1]+"-"+info[2]+"-"+info[3]+" "+info[6][:2]+":"+info[6][2:4])
	if terr1 != nil {
		color.Set(color.FgHiRed)
		fmt.Println(terr1)
		color.Unset()
		return 0
	}
	if terr2 != nil {
		color.Set(color.FgHiRed)
		fmt.Println(terr2)
		color.Unset()
		return 0
	}
	diff := t2.Sub(t1)
	breaktime, breakTimeParseErr := strconv.Atoi(info[8])
	var adjusted int
	if breakTimeParseErr != nil {
		adjusted = int(diff.Minutes())
	} else {
		adjusted = int(diff.Minutes()) - breaktime
	}
	color.Set(color.FgHiBlack)
	fmt.Println("\t", t1.Format("2006-01-02"), ":", diff, ":", adjusted, "minutes")
	color.Unset()
	if adjusted < 0 {
		return 0
	}
	return adjusted
}
