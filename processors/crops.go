package processors

import (
	"github.com/jaysonesmith/stardew-crops/data"

	"github.com/jaysonesmith/stardew-crops/output"
)

var cropData data.CropData

// AllCrops returns all crop data raw
func AllCrops(args ...string) {
	output.Print(cropData)
}

func init() {
	cropData = data.InitData()
}
