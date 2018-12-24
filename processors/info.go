package processors

import (
	"fmt"

	"github.com/jaysonesmith/stardew-crops/data"
)

// Info returns the data around the specified crop
func Info(args ...string) (data.CropData, error) {
	if len(args) == 0 || args[0] == "" {
		return data.CropData{}, fmt.Errorf("No crop name specified, please try again")
	}

	for _, crop := range cropData.Crops {
		if args[0] == crop.Name {
			return data.CropData{Crops: []data.Crop{crop}}, nil
		}
	}

	return data.CropData{}, fmt.Errorf("Unable to find matching crop for %s", args[0])
}

func init() {
	cropData = data.InitData()
}
