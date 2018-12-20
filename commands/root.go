package commands

import (
	"github.com/spf13/cobra"
)

// StardewCrops represents the base command when called without any subcommands
var StardewCrops = &cobra.Command{
	Use:   "stardew-crops",
	Short: "Information around Stardew's Crops",
	Long:  `A cli tool for information around Stardew Valley's Crops`,
}

func init() {
	StardewCrops.PersistentFlags().StringVarP(&format, "format", "f", "pretty", "set desired output type. choices: pretty, json, raw")
}
