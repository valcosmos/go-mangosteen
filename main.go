package main

import (
	"mangosteen/cmd"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	cmd.Run()
}
