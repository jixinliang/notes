package main

import (
	"flag"
	"fmt"
	"strings"
)

type Flags struct {
	Names []string
}

func (f *Flags) GetNames() []string {
	return f.Names
}

func (f *Flags) String() string {
	return fmt.Sprint(f.Names)
}

func (f *Flags) Set(value string) error {
	if len(f.Names) > 0 {
		return fmt.Errorf("Cannot use named flags more than once!")
	}

	names := strings.Split(value, ",")
	for _, item := range names {
		f.Names = append(f.Names, item)
	}
	return nil
}

func main() {
	var names Flags
	k := flag.Int("k", 0, "An int")
	o := flag.String("o", "Jack", "The name")
	flag.Var(&names, "names", "Comma-separated list")
	flag.Parse()

	fmt.Println("-k:", *k)
	fmt.Println("-o:", *o)

	for i, i2 := range names.GetNames() {
		fmt.Println(i, i2)
	}

	fmt.Println("Remained command line arguments:")
	for idx, val := range flag.Args() {
		fmt.Println(idx, ":", val)
	}
}

//  go run advFlag.go -k 10 -o=ok -names=jay,jack 10 one two