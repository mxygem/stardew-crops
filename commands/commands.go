package commands

import (
	"fmt"
	"os"
)

var cropData map[string]interface{}

var cfgFile string
var format string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := StardewCrops.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
