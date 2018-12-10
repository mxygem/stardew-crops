package commands

import (
	"github.com/jaysonesmith/stardew-crops/processors"
	"github.com/spf13/cobra"
)

// infoCmd returns data based on the input crop
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get information on a specific crop",
	Long:  `Specify a crop name to get more information on it i.e. 'stardew-crops info garlic'`,
	Run:   info,
}

func info(cmd *cobra.Command, args []string) {
	processors.Info(args...)
}

func init() {
	stardewCrops.AddCommand(infoCmd)
}
