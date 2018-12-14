package processors

import (
	"fmt"

	"github.com/jaysonesmith/stardew-crops/data"
)

// Info returns the data around the specified crop
func Info(args ...string) (data.Crop, error) {
	if len(args) == 0 || args[0] == "" {
		return data.Crop{}, fmt.Errorf("No crop name specified, please try again")
	}

	for _, crop := range cropData.Crops {
		if args[0] == crop.Name {
			return crop, nil
		}
	}

	return data.Crop{}, fmt.Errorf("Unable to find matching crop for %s", args[0])
}

func init() {
	cropData = data.InitData()
}
