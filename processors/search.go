package processors

import (
	"fmt"
	"strconv"

	"github.com/jaysonesmith/stardew-crops/data"
)

// Search ...
func Search(flags map[string]string) ([]data.Crop, error) {
	var out []data.Crop

	c := cropData.Crops
	for i := 0; i < len(cropData.Crops); i++ {
		specified, ok, err := byBundle(c[i], flags)
		if err != nil {
			return []data.Crop{}, err
		} else if specified && !ok {
			continue
		}

		specified, ok, err = byGrowth("g", c[i].Info.GrowthTime, flags)
		if err != nil {
			return []data.Crop{}, err
		} else if specified && !ok {
			continue
		}

		specified, ok, err = byGrowth("l", c[i].Info.GrowthTime, flags)
		if err != nil {
			return []data.Crop{}, err
		} else if specified && !ok {
			continue
		}

		specified, ok, err = bySeason(c[i].Info.Season, flags)
		if err != nil {
			return []data.Crop{}, err
		} else if specified && !ok {
			continue
		}

		specified, ok, err = byTrellis(c[i].Info.Trellis, flags)
		if err != nil {
			return []data.Crop{}, err
		} else if specified && !ok {
			continue
		}

		specified, ok, err = byContinuous(c[i].Info.Continual, flags)
		if err != nil {
			return []data.Crop{}, err
		} else if specified && !ok {
			continue
		}

		out = append(out, c[i])
	}

	if len(out) == 0 {
		return []data.Crop{}, fmt.Errorf("No matching crops found")
	}

	return out, nil
}

func byBundle(c data.Crop, f map[string]string) (bool, bool, error) {
	v, ok := f["bundle"]
	if !ok {
		return false, false, nil
	} else if ok && v == "" {
		return true, false, fmt.Errorf(valueRequired("bundle"))
	}

	for _, bundle := range c.Bundles {
		if v == bundle {
			return true, true, nil
		}
	}
	return true, false, nil
}

func byGrowth(o string, t int64, f map[string]string) (bool, bool, error) {
	if o == "g" {
		v, ok := f["growthgt"]
		if !ok {
			return false, false, nil
		} else if ok && v == "" {
			return false, false, fmt.Errorf(valueRequired("growthgt"))
		}

		vi, _ := strconv.Atoi(v)
		if t >= int64(vi) {
			return false, true, nil
		}
	}

	if o == "l" {
		v, ok := f["growthlt"]
		if !ok {
			return false, false, nil
		} else if ok && v == "" {
			return false, false, fmt.Errorf(valueRequired("growthlt"))
		}

		vi, _ := strconv.Atoi(v)
		if t <= int64(vi) {
			return false, true, nil
		}
	}

	return true, false, nil
}

func bySeason(cs []string, f map[string]string) (bool, bool, error) {
	v, ok := f["season"]
	if !ok {
		return false, false, nil
	} else if ok && v == "" {
		return false, false, fmt.Errorf(valueRequired("season"))
	}

	for _, season := range cs {
		if v == season {
			return true, true, nil
		}
	}

	return true, false, nil
}

func byTrellis(t bool, f map[string]string) (bool, bool, error) {
	v, ok := f["trellis"]
	if !ok {
		return false, false, nil
	} else if ok && v == "" {
		return false, false, fmt.Errorf(valueRequired("trellis"))
	}

	parsedV, err := strconv.ParseBool(v)
	if err != nil {
		return true, false, err
	}

	if t == parsedV {
		return true, true, nil
	}

	return true, false, nil
}

func byContinuous(c bool, f map[string]string) (bool, bool, error) {
	v, ok := f["continuous"]
	if !ok {
		return false, false, nil
	} else if ok && v == "" {
		return false, false, fmt.Errorf(valueRequired("continuous"))
	}

	parsedV, err := strconv.ParseBool(v)
	if err != nil {
		return true, false, err
	}

	if c == parsedV {
		return true, true, nil
	}

	return true, false, nil
}

func valueRequired(flag string) string {
	return fmt.Sprintf("A value must be provided for the %s flag", flag)
}

func init() {
	cropData = data.InitData()
}
