package processors

import (
	"fmt"

	"github.com/jaysonesmith/stardew-crops/data"

	"github.com/jaysonesmith/stardew-crops/output"
)

// Season ...
func Season(args ...string) {
	cbs := CropsByAllSeasons()
	fmt.Println(cbs)

	if len(args) == 0 || args[0] == "" {
		output.Print(fmt.Sprintf("No crop name specified, please try again"))
		return
	}

	for _, crop := range cropData.Crops {
		if args[0] == crop.Name {
			output.Print(crop)
			return
		}
	}

	output.Print(fmt.Sprintf("Unable to find matching crop for %s", args[0]))
}

// CropsByAllSeasons returns all crops grouped by season
func CropsByAllSeasons() map[string][]string {
	return map[string][]string{}
}

func init() {
	cropData = data.InitData()
}
