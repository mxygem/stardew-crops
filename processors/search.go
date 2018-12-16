package processors

import (
	"fmt"

	"github.com/jaysonesmith/stardew-crops/data"
)

// Search ...
func Search(flags map[string]string) ([]data.Crop, error) {
	var out []data.Crop

	for k, v := range flags {
		switch k {
		case "bundle":
			if v == "" {
				return []data.Crop{}, fmt.Errorf(valueRequired("bundle"))
			}
			out = append(out, byBundle(v)...)
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

	if len(out) == 0 {
		return []data.Crop{}, fmt.Errorf("No matching crops found")
	}

	return out, nil
}

func byBundle(v string) []data.Crop {
	var matched []data.Crop
	for _, c := range cropData.Crops {
		for _, b := range c.Bundles {
			if b == v {
				matched = append(matched, c)
			}
		}
	}

	return matched
}

func valueRequired(flag string) string {
	return fmt.Sprintf("A value must be provided for the %s flag", flag)
}

func init() {
	cropData = data.InitData()
}
