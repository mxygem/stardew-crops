package commands

import (
	"github.com/jaysonesmith/stardew-crops/output"
	"github.com/jaysonesmith/stardew-crops/processors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// searchCmd returns data based on the input crop
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for crops via a number of methods",
	Long:  "Search returns a list of crops based on the provided filters and search criteria",
	Run:   Search,
}

// SearchFlags is a collection of the supported
// flags for the search command
type SearchFlags struct {
	Bundle     string
	Continuous bool
	GrowthGT   int64
	GrowthLT   int64
	Season     string
	Trellis    bool
}

var sf SearchFlags

// Search is the main orchestrator for the cli's
// search functionality
func Search(cmd *cobra.Command, args []string) {
	// return help information if no flags are provided
	if !userSetFlags(cmd) {
		cmd.Help()
		return
	}

	out, err := processors.Search(args...)
	if err != nil {
		output.Print(err.Error())
		return
	}

	output.Print(out)
}

// userSetFlags checks for flags that were
// explicitly set to determine if any were set by
// the user
func userSetFlags(cmd *cobra.Command) (found bool) {
	found = false
	cmd.Flags().Visit(func(f *pflag.Flag) {
		if f.Changed {
			found = true
		}
	})

	return
}

func init() {
	// set flags
	searchCmd.Flags().StringVarP(&sf.Bundle, "bundle", "b", "", "Specify a bundle name to search for")
	searchCmd.Flags().BoolVarP(&sf.Continuous, "continuous", "c", false, "Search for crops that grow continuously")
	searchCmd.Flags().Int64VarP(&sf.GrowthGT, "growthgt", "g", 0, "Search for crops that take the specified time or longer to grow")
	searchCmd.Flags().Int64VarP(&sf.GrowthLT, "growthlt", "l", 0, "Search for crops that take the specified time or less to grow")
	searchCmd.Flags().StringVarP(&sf.Season, "season", "s", "", "Search by a season. Options: Spring, Summer, Fall")
	searchCmd.Flags().BoolVarP(&sf.Trellis, "trellis", "t", false, "Search for crops. Options: Spring, Summer, Fall")

	stardewCrops.AddCommand(searchCmd)
}
