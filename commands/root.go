package commands

import (
	"github.com/spf13/cobra"
)

// stardewCrops represents the base command when called without any subcommands
var stardewCrops = &cobra.Command{
	Use:   "stardew-crops",
	Short: "Information around Stardew's Crops",
	Long:  `A cli tool for information around Stardew Valley's Crops`,
}

func init() {
	stardewCrops.PersistentFlags().StringVarP(&format, "format", "f", "pretty", "set desired output type. choices: pretty, json")
}