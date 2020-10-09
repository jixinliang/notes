package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("json")
	viper.SetConfigFile("myConfig.json")
	fmt.Println("Using config:", viper.ConfigFileUsed())
	viper.ReadInConfig()

	if viper.IsSet("item1.key3") {
		fmt.Println("item1.key3", viper.Get("item1.key3"))
	} else {
		fmt.Println("item1.key1 not set")
	}

	if viper.IsSet("item2.key2") {
		fmt.Println("item2.key2", viper.Get("item2.key2"))
	} else {
		fmt.Println("item2.key2 not set")
	}
}
