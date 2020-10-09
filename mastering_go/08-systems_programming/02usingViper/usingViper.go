package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.BindEnv("GOMAXPROCS")
	val := viper.Get("GOMAXPROCS")
	fmt.Println("go max procs:", val)

	viper.Set("GOMAXPROCS", 10)
	val = viper.Get("GOMAXPROCS")
	fmt.Println("go max procs after costom:", val)

	viper.BindEnv("NEW_VARIABLE")
	val = viper.Get("NEW_VARIABLE")
	if val == nil {
		fmt.Println("new variable not defined.")
		return
	}
	fmt.Println(val)
}
