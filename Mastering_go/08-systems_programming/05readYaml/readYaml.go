package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func main() {
	var configFile *string
	configFile = flag.String("c", "myConfig", "Setting my configuration file")
	flag.Parse()

	_, err := os.Stat(*configFile)
	if err != nil {
		viper.SetConfigName(*configFile)
		viper.AddConfigPath("/tmp")
		viper.AddConfigPath(".")
	} else {
		fmt.Println("Using use specified configuration file!")
		viper.SetConfigFile(*configFile)
	}

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("read config file failed", err)
		return
	}

	if viper.IsSet("item1.k1") {
		fmt.Println("item1.val1:", viper.Get("item1.k1"))
	} else {
		fmt.Println("item1.k1 not set!")
	}
	if viper.IsSet("item1.k2") {
		fmt.Println("item1.val2:", viper.Get("item1.k2"))
	} else {
		fmt.Println("item1.k2 not set!")
	}
	if !viper.IsSet("item3.k1") {
		fmt.Println("item3.k1 is not set!")
	}
}
