package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jaysonesmith/stardew-crops/utils"
)

// InitData opens the appropriate data file and returns it
func InitData() CropData {
	fileBytes := utils.OpenBytes(cropsDataFile())

	var cropData CropData
	if err := json.Unmarshal(fileBytes, &cropData); err != nil {
		fmt.Println("unable to process crops file:", err)
		os.Exit(1)
	}

	return cropData
}

func cropsDataFile() string {
	// if os.Getenv("TEST") == "true" {
	// 	return "./test_data/crops.json"
	// }
	return "./crops.json"
}
