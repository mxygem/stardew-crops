package processors

import (
	"fmt"

	"github.com/jaysonesmith/stardew-crops/data"
)

// Search ...
func Search(flags map[string]interface{}) ([]string, error) {
	for k, v := range flags {
		switch k {
		case "bundle":
			if v == "Summer Crops" {
				return []string{"blueberry", "hot pepper"}, nil
			} else if v == "Quality Crops" {
				return []string{"corn"}, nil
			}
		case "continuous":
			fmt.Println("continuous value:", v)
		case "growthgt":
			fmt.Println("growthgt value:", v)
		case "growthlt":
			fmt.Println("growthlt value:", v)
		case "season":
			fmt.Println("season value:", v)
		case "trellis":
			fmt.Println("trellis value:", v)
		}
	}

	// for _, crop := range cropData.Crops {
	// 	if args[0] == crop.Name {
	// 		return crop, nil
	// 	}
	// }

	return []string{}, nil
}

func init() {
	cropData = data.InitData()
}
