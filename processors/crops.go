package processors

import (
	"github.com/jaysonesmith/stardew-crops/data"
)

var cropData data.CropData

// AllCrops returns all crop data raw
func AllCrops(args ...string) data.CropData {
	return cropData
}

func init() {
	cropData = data.InitData()
}
