package commands

import (
	"github.com/jaysonesmith/stardew-crops/output"
	"github.com/jaysonesmith/stardew-crops/processors"
	"github.com/spf13/cobra"
)

// searchCmd returns data based on the input crop
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for crops via a number of methods",
	Long:  "Search returns a list of crops based on the provided filters and search criteria",
	Run:   Search,
}

// Search is the main orchestrator for the cli's search functionality
func Search(cmd *cobra.Command, args []string) {
	out, err := processors.Search(args...)
	if err != nil {
		output.Print(err.Error())
		return
	}

	output.Print(out)
}

func init() {
	stardewCrops.AddCommand(infoCmd)
}
