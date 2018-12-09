package processors

import (
	"fmt"

	"github.com/jaysonesmith/stardew-crops/data"

	"github.com/jaysonesmith/stardew-crops/output"
)

var cropData map[string]interface{}

// All returns all crop data raw
func All(args ...string) {
	fmt.Println("cropData:", cropData)
	for i, arg := range args {
		fmt.Printf("%d: %s\n", i, arg)
	}

	output.Print(cropData)
}

func init() {
	cropData = data.InitData()
}
