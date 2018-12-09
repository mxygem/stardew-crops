package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// InitData opens the appropriate data file and returns it
func InitData() CropData {
	fileBytes, err := ioutil.ReadFile(cropsDataFile())
	if err != nil {
		fmt.Println("unable to open crops file:", err)
		os.Exit(1)
	}

	var cropData CropData
	if err := json.Unmarshal(fileBytes, &cropData); err != nil {
		fmt.Println("unable to process crops file:", err)
		os.Exit(1)
	}

	return cropData
}

func cropsDataFile() string {
	if os.Getenv("TEST") == "true" {
		return "./test_data/crops.json"
	}
	return "./data/crops.json"
}
