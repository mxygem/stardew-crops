package processors

import (
	"github.com/jaysonesmith/stardew-crops/data"

	"github.com/jaysonesmith/stardew-crops/output"
)

// Info returns the data around the specified crop
func Info(args ...string) {
	// for i, arg := range args {
	// 	fmt.Printf("%d: %s\n", i, arg)
	// }

	output.Print(cropData)
}

func init() {
	cropData = data.InitData()
}
