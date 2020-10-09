package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: go run %s filename\n", filepath.Base(args[0]))
		os.Exit(1)
	}

	fileName := args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Open file failed:", err)
		os.Exit(2)
	}
	defer file.Close()

	notMatch := 0
	reader := bufio.NewReader(file)

	for {
		lineStr, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Read file failed:", err)
		}

		reg1 := regexp.MustCompile(`.*\[(\d{2}\/\w+/\d{4}:\d{2}:\d{2}:\d{2}.*)\].*`)
		//reg1 := regexp.MustCompile(`.*\[(\d\d\/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\].*`)
		if reg1.MatchString(lineStr) {
			matchedStr := reg1.FindStringSubmatch(lineStr)
			//fmt.Println("matchedStr:----------", matchedStr)
			date1, err := time.Parse("02/Jan/2006:15:04:05 -0700", matchedStr[1])
			if err != nil {
				notMatch++
			} else {
				newFormat := date1.Format(time.Stamp)
				fmt.Print(strings.Replace(lineStr, matchedStr[1], newFormat, 1))
			}
			continue
		}

		//reg2 := regexp.MustCompile(`.*\[(\w+\-\d\d-\d\d:\d\d:\d\d:\d\d.*)\].*`)
		reg2 := regexp.MustCompile(`.*\[(\w+\-\d{2}-\d{2}:\d{2}:\d{2}:\d{2}.*)\].*`)
		if reg2.MatchString(lineStr) {
			matchedStr := reg2.FindStringSubmatch(lineStr)
			date1, err := time.Parse("Jan-02-06:15:04:05 -0700", matchedStr[1])
			if err != nil {
				notMatch++
			} else {
				newFormat := date1.Format(time.Stamp)
				fmt.Print(strings.Replace(lineStr, matchedStr[1], newFormat, 1))
			}
			continue
		}
	}
	fmt.Println(notMatch, "Lines not matched!")
}
