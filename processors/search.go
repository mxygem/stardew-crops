package processors

import (
	"fmt"
	"strconv"

	"github.com/jaysonesmith/stardew-crops/data"
)

// Search ...
func Search(flags map[string]string) ([]data.Crop, error) {
	var out []data.Crop

	for k, v := range flags {
		// TODO: manage upper/lowercases
		// v = strings.ToLower(v)
		switch k {
		case "bundle":
			if v == "" {
				return []data.Crop{}, fmt.Errorf(valueRequired(k))
			}
			out = append(out, byBundle(v)...)
		case "continuous":
			fmt.Println("continuous value:", v)
		case "growthgt":
			if v == "" {
				return []data.Crop{}, fmt.Errorf(valueRequired(k))
			}
			out = append(out, byGrowth("g", v)...)
		case "growthlt":
			if v == "" {
				return []data.Crop{}, fmt.Errorf(valueRequired(k))
			}
			out = append(out, byGrowth("l", v)...)
		case "season":
			if v == "" {
				return []data.Crop{}, fmt.Errorf(valueRequired(k))
			}
			out = append(out, bySeason(v)...)
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

func byGrowth(o, v string) []data.Crop {
	var matched []data.Crop
	vi, _ := strconv.Atoi(v)
	for _, c := range cropData.Crops {
		if o == "g" && c.Info.GrowthTime >= int64(vi) {
			matched = append(matched, c)
		} else if o == "l" && c.Info.GrowthTime <= int64(vi) {
			matched = append(matched, c)
		}
	}

	return matched
}

func bySeason(s string) []data.Crop {
	var matched []data.Crop
	for _, c := range cropData.Crops {
		for _, b := range c.Info.Season {
			if b == s {
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
