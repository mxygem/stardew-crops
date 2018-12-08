package processors

import (
	"fmt"

	"github.com/jaysonesmith/stardew-crops/output"
)

// All returns all crop data raw
func All(args []string, cropData *map[string]interface{}) {

	fmt.Println("cropData:", cropData)
	for i, arg := range args {
		fmt.Printf("%d: %s\n", i, arg)
	}

	output.Print(cropData)
}
