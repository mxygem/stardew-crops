package processors

import (
	"fmt"

	"github.com/jaysonesmith/stardew-crops/data"
)

// Search ...
func Search(flags map[string]interface{}) ([]string, error) {
	var out []string
	for k, v := range flags {
		switch k {
		case "bundle":
			out = append(out, byBundle(v.(string))...)
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
		return []string{}, fmt.Errorf("no matching crops found")
	}

	return out, nil
}

func byBundle(v string) []string {
	var matched []string
	for _, c := range cropData.Crops {
		for _, b := range c.Bundles {
			if b == v {
				matched = append(matched, c.Name)
			}
		}
	}

	return matched
}

func init() {
	cropData = data.InitData()
}
