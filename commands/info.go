package commands

import (
	"github.com/jaysonesmith/stardew-crops/output"
	"github.com/jaysonesmith/stardew-crops/processors"
	"github.com/spf13/cobra"
)

// infoCmd returns data based on the input crop
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get information on a specific crop",
	Long:  `Specify a crop name to get more information on it i.e. 'stardew-crops info garlic'`,
	Run:   Info,
}

func Info(cmd *cobra.Command, args []string) {
	out, err := processors.Info(args...)
	if err != nil {
		output.Print(err.Error())
		return
	}

	output.Print(out)
}

func init() {
	stardewCrops.AddCommand(infoCmd)
}
