package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Standard output is more or less equivalent to printing on the screen.

// Using standard output
func stdOutput() {
	var myStr string
	args := os.Args

	if len(args) == 1 {
		myStr = "Pls give me one or more arguments!"
		return
	}
	myStr = args[1]

	_, _ = io.WriteString(os.Stdout, myStr)
	_, _ = io.WriteString(os.Stdout, "\n")
}

// Reading from standard input
func stdInput() {
	var file *os.File
	file = os.Stdin
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(":", scanner.Text())
	}
}

// Working with command-line arguments
func cmdLineArgs() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Pls give me one or more 'Floats'")
		os.Exit(1)
	}

	min, _ := strconv.ParseFloat(args[1], 64)
	max, _ := strconv.ParseFloat(args[1], 64)

	for i := 2; i < len(args); i++ {
		n, _ := strconv.ParseFloat(args[i], 64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("min:", min)
	fmt.Println("max:", max)
}

// error output
func stdErrOutput() {
	var myStr string
	args := os.Args

	if len(args) == 1 {
		myStr = "Pls give me one or more argumts!"
		return
	}
	myStr = args[1]

	_, _ = io.WriteString(os.Stdout, "This is standard out\n")
	_, _ = io.WriteString(os.Stderr, myStr)
	_, _ = io.WriteString(os.Stderr, "\n")

}

// Writing to a custom log file
func customLog() {
	myLogFile := "/tmp/customLog.log"
	file, err := os.OpenFile(myLogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	myLog := log.New(file, "customLogLineNum: ", log.LstdFlags)
	myLog.SetFlags(log.LstdFlags | log.Lshortfile)
	myLog.Println("line number 1")
	myLog.Println("line number 2")
}

// The error data type
func customErr(a, b int) error {
	if a == b {
		err := errors.New("'Error' in customError function")
		return err
	} else {
		return nil
	}
}

func cmdLineArgs1() {
	args := os.Args
	k := 1
	err := errors.New("a New error")
	var n float64

	if len(args) == 1 {
		fmt.Println("Pls give me one or more Floats")
		os.Exit(1)
	}

	for err != nil {
		if k > len(args) {
			fmt.Println("none of the arguments is Float")
			return
		}

		n, err = strconv.ParseFloat(args[1], 64)
		k++
	}
	min, max := n, n

	for i := 2; i < len(args); i++ {
		n, err = strconv.ParseFloat(args[i], 64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println("min: ", min)
	fmt.Println("max: ", max)
}

func main() {
	//err := customErr(1, 1)
	//if err != nil {
	//	//fmt.Println(err)
	//	panic(err)
	//} else {
	//	fmt.Println("Ok!~")
	//}

	cmdLineArgs1()
}

/*
There are three main ways to get user input:
firstly, by reading the command-line arguments of a program;
secondly, by asking the user for input;
or thirdly, by reading external files.


calling os.Exit() from a function other than main() is considered a bad practice.
Functions other than main() tend to return the error message before exiting,
*/
