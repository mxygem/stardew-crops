package commands

import (
	"github.com/jaysonesmith/stardew-crops/output"
	"github.com/jaysonesmith/stardew-crops/processors"
	"github.com/spf13/cobra"
)

// seasonCmd returns data based on the input crop
var seasonCmd = &cobra.Command{
	Use:   "season",
	Short: "List crops by season or filter to a specific one",
	Long:  `Specify a season to limit to only that season i.e. 'stardew-crops season summer'`,
	Run:   season,
}

func season(cmd *cobra.Command, args []string) {
	out := processors.Season(args...)
	output.Print(out)
}

func init() {
	stardewCrops.AddCommand(seasonCmd)
}
