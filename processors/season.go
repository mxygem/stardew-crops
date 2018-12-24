package processors

// import (
// 	"fmt"

// 	"github.com/jaysonesmith/stardew-crops/data"
// )

// type CropsBySeasons struct {
// 	Spring []string `json:"spring,omitempty"`
// 	Summer []string `json:"summer,omitempty"`
// 	Fall   []string `json:"fall,omitempty"`
// }

// // Season ...
// func Season(args ...string) (CropsBySeasons, error) {
// 	cbs := cropsByAllSeasons()

// 	if len(args) == 0 || args[0] == "" {
// 		return cbs, nil
// 	} else if args[0] == "spring" {
// 		return CropsBySeasons{Spring: cbs.Spring}, nil
// 	} else if args[0] == "summer" {
// 		return CropsBySeasons{Summer: cbs.Summer}, nil
// 	} else if args[0] == "fall" {
// 		return CropsBySeasons{Fall: cbs.Fall}, nil
// 	}

// 	return CropsBySeasons{}, fmt.Errorf("Unknown season for %s", args[0])
// }

// func cropsByAllSeasons() CropsBySeasons {
// 	out := CropsBySeasons{}
// 	for _, crop := range cropData.Crops {
// 		for _, season := range crop.Info.Season {
// 			switch season {
// 			case "spring":
// 				out.Spring = append(out.Spring, crop.Name)
// 			case "summer":
// 				out.Summer = append(out.Summer, crop.Name)
// 			case "fall":
// 				out.Fall = append(out.Fall, crop.Name)
// 			}
// 		}
// 	}

// 	return out
// }

// func init() {
// 	cropData = data.InitData()
// }
