package processors

import (
	"fmt"

	"github.com/jaysonesmith/stardew-crops/data"
	"github.com/jaysonesmith/stardew-crops/output"
)

// Season ...
func Season(args ...string) CropsBySeasons {
	cbs := cropsByAllSeasons()
	// fmt.Println(cbs)

	if len(args) == 0 || args[0] == "" {
		return cbs
	} else if args[0] == "spring" {
		return CropsBySeasons{Spring: cbs.Spring}
	} else if args[0] == "summer" {
		return CropsBySeasons{Summer: cbs.Summer}
	} else if args[0] == "fall" {
		return CropsBySeasons{Fall: cbs.Fall}
	}

	output.Print(fmt.Sprintf("Unknown season for %s", args[0]))
	return CropsBySeasons{}
}

type CropsBySeasons struct {
	Spring []string `json:"spring,omitempty"`
	Summer []string `json:"summer,omitempty"`
	Fall   []string `json:"fall,omitempty"`
}

func cropsByAllSeasons() CropsBySeasons {
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
