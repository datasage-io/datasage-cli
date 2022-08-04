package main

import (
	"flag"
	"fmt"

	"github.com/datasage-io/datasage-cli/cmd"
	"github.com/spf13/viper"
)

//Read Configuration File
var configFilePath *string

func main() {
	configFilePath = flag.String("config-path", "configuration/", "configuration/")
	flag.Parse()

	loadConfig()
	//Execute Commands
	cmd.Execute()
}

// loadConfig - Load the config parameters
func loadConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(*configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		if readErr, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("No config file found at %s\n", *configFilePath)
		} else {
			fmt.Printf("Error reading config file: %s\n", readErr)
		}
	}
}
