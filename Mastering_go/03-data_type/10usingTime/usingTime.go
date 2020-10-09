package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

func times() {
	fmt.Println("Epoch time:", time.Now().Unix())

	nowTime := time.Now()
	fmt.Println(nowTime, "->>", nowTime.Format(time.RFC3339))
	fmt.Println(nowTime.Year(), nowTime.Month(), nowTime.Hour(), nowTime.Minute())

	time.Sleep(time.Second)
	t1 := time.Now()
	fmt.Println("Duration:", t1.Sub(nowTime))
}

func timeFormat() {
	nowTime := time.Now()
	timeFormat := nowTime.Format("01 January 2006")
	fmt.Println(timeFormat)
	loc, err := time.LoadLocation("Europe/London")
	if err != nil {
		fmt.Println(err)
	}
	londonTime := nowTime.In(loc)
	fmt.Println("London time:", londonTime)
}

func parseHour() {
	var myTime string
	args := os.Args

	if len(args) != 2 {
		fmt.Printf("Usage: %s string\n", filepath.Base(args[0]))
		os.Exit(1)
	}
	myTime = args[1]

	hourTime, err := time.Parse("15:04", myTime)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Full:", hourTime)
		fmt.Println("Time:", hourTime.Hour(), hourTime.Minute())
	}
}

func parseFull() {
	var myTime string
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %s string\n", filepath.Base(args[0]))
		os.Exit(1)
	}
	myTime = args[1]

	yearTime, err := time.Parse("02 January 2006", myTime)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Full:", yearTime)
		fmt.Println("Time:", yearTime.Day(), yearTime.Month(), yearTime.Year())
	}
}

func main() {
	logs := []string{"127.0.0.1 - - [16/Nov/2020:10:49:46 +0200]325504",
		"127.0.0.1 - - [16/Nov/2020:10:16:41 +0200] \"GET /CVEN HTTP/1.1\" 200 12531 \"-\" \"Mozilla/5.0 AppleWebKit/537.36",
		"127.0.0.1 200 9412 - - [12/Nov/2020:06:26:05 +0200] \"GET \"http://www.mtsoukalos.eu/taxonomy/term/47\" 1507",
		"[12/Nov/2020:16:27:21 +0300]",
		"[12/Nov/2020:20:88:21 +0200]",
		"[12/Nov/2020:20:21 +0200]",
	}

	for _, logEntry := range logs {
		reg := regexp.MustCompile(`^.*\[(\d\d\/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\].*$`)
		if reg.MatchString(logEntry) {
			matched := reg.FindStringSubmatch(logEntry)
			dt, err := time.Parse("02/Jan/2006:15:04:05 -0700", matched[1])
			if err != nil {
				fmt.Println("Not valied data time format")
			} else {
				newFormat := dt.Format(time.RFC850)
				fmt.Println(newFormat)
			}
		} else {
			fmt.Println("Not matched")
		}
	}

}
