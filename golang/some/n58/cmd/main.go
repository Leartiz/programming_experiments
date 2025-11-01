package main

import (
	"fmt"

	"github.com/integrii/flaggy"
)

func main() {
	var configFile string = "config_example.yaml"
	var boolFlagB bool

	flaggy.DefaultParser.Name = "EIR Webhook Notification"
	flaggy.DisableCompletion()

	flaggy.String(&configFile, "c", "config", "Config file path")
	flaggy.Bool(&boolFlagB, "y", "yes", "A sample boolean flag")

	flaggy.Parse()
	fmt.Println(configFile)

	fmt.Println("OK")
}
