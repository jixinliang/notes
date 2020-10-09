package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	myStr := "Hello World"

	fmt.Printf("to upper: %s\n", strings.ToUpper(myStr))
	fmt.Println("to lower:", strings.ToLower(myStr))
	fmt.Println("to title:", strings.ToTitle(myStr))
	myStr1 := "hello, world"
	fmt.Println("equalFold:", strings.EqualFold(myStr, myStr1))

	fmt.Println("hasPrefix:", strings.HasPrefix(myStr, "h"))
	fmt.Println("hasSuffix:", strings.HasSuffix(myStr, "d"))

	fmt.Println("counts:", strings.Count(myStr, "o"))

	fmt.Println("repeat:", strings.Repeat(myStr, 3))

	fmt.Println("trimSpace:", strings.TrimSpace("    hi   "))
	fmt.Println("trimLeft:", strings.TrimLeft("\t\thhh", "\n\t"))
	fmt.Println("trimRight", strings.TrimRight("hh\n\t", "\t\n"))

	fmt.Println("compare:", strings.Compare(myStr, myStr1))

	fmt.Println("fields:", strings.Fields("fields test"))

	fmt.Println("split:", strings.Split(myStr1, ", "))
	fmt.Println("splitAfter:", strings.SplitAfter("a++,b++", ","))

	fmt.Println("replace:", strings.Replace("abcd efg", "", "_", -1))

	fmt.Println("join:", strings.Join([]string{"Hello", "World"}, "-"))

	isLetter := func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsNumber(r) }
	fmt.Println(strings.TrimFunc("ad 123 ! \t", isLetter))

}
