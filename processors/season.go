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

type CropsBySeasons struct {
	Spring []string `json:"spring,omitempty"`
	Summer []string `json:"summer,omitempty"`
	Fall   []string `json:"fall,omitempty"`
}

// CropsByAllSeasons returns all crops grouped by season
func CropsByAllSeasons() CropsBySeasons {
	out := CropsBySeasons{}
	for _, crop := range cropData.Crops {
		switch crop.Info.Season {
		case "spring":
			out.Spring = append(out.Spring, crop.Name)
		case "summer":
			out.Summer = append(out.Summer, crop.Name)
		case "fall":
			out.Fall = append(out.Fall, crop.Name)
		}
	}

	return out
}

func init() {
	cropData = data.InitData()
}
