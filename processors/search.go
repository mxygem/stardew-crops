package processors

import (
	"github.com/jaysonesmith/stardew-crops/data"
)

// Search ...
func Search(args ...string) ([]string, error) {
	// if len(args) == 0 || args[0] == "" {
	// 	return data.Crop{}, fmt.Errorf("No crop name specified, please try again")
	// }

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
