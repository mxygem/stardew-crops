package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// infoCmd returns data based on the input crop
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get information on a specific crop",
	Long:  `Specify a crop name to get more information on it i.e. 'stardew-crops info garlic'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cropData:", &cropData)
	},
}

func init() {
	stardewCrops.AddCommand(infoCmd)
}
