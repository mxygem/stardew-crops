package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var cropData map[string]interface{}

var cfgFile string
var format string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := stardewCrops.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initDB() {
	fileBytes, err := ioutil.ReadFile("./data/crops.json")
	if err != nil {
		fmt.Println("unable to open crops file:", err)
		os.Exit(1)
	}

	if err := json.Unmarshal(fileBytes, &cropData); err != nil {
		fmt.Println("unable to process crops file:", err)
		os.Exit(1)
	}
}
