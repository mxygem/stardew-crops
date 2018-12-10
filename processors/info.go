package processors

import (
	"fmt"

	"github.com/jaysonesmith/stardew-crops/data"

	"github.com/jaysonesmith/stardew-crops/output"
)

// Info returns the data around the specified crop
func Info(args ...string) {
	if len(args) == 0 || args[0] == "" {
		output.Print(fmt.Sprintf("No crop name specified, please try again"))
		return
	}

	for _, crop := range cropData.Crops {
		if args[0] == crop.Name {
			output.Print(crop)
			return
		}
	}

	output.Print(fmt.Sprintf("Unable to find matching crop for %s", args[0]))
}

func init() {
	cropData = data.InitData()
}
