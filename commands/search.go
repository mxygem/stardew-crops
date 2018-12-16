package commands

import (
	"github.com/jaysonesmith/stardew-crops/output"
	"github.com/jaysonesmith/stardew-crops/processors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// SearchCmd returns data based on the input crop
var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for crops via a number of methods",
	Long:  "Search returns a list of crops based on the provided filters and search criteria",
	Run:   Search,
}

// Search is the main orchestrator for the cli's
// search functionality
func Search(cmd *cobra.Command, args []string) {
	userFlags := userSetFlags(cmd)

	// return help information if no flags are provided
	if len(userFlags) == 0 {
		cmd.Help()
		return
	}

	out, err := processors.Search(userFlags)
	if err != nil {
		output.Print(err.Error())
		return
	}

	output.Print(out)
}

// userSetFlags checks for flags that were
// explicitly set to determine if any were set by
// the user
func userSetFlags(cmd *cobra.Command) map[string]string {
	userFlags := make(map[string]string, 0)
	cmd.Flags().Visit(func(f *pflag.Flag) {
		if f.Changed {
			userFlags[f.Name] = f.Value.String()
		}
	})

	return userFlags
}

func init() {
	// set flags
	SearchCmd.Flags().StringP("bundle", "b", "", "Specify a bundle name to search for")
	SearchCmd.Flags().BoolP("continuous", "c", false, "Search for crops that grow continuously")
	SearchCmd.Flags().Int64P("growthgt", "g", 0, "Search for crops that take the specified time or longer to grow")
	SearchCmd.Flags().Int64P("growthlt", "l", 0, "Search for crops that take the specified time or less to grow")
	SearchCmd.Flags().StringP("season", "s", "", "Search by a season. Options: Spring, Summer, Fall")
	SearchCmd.Flags().BoolP("trellis", "t", false, "Search for crops. Options: Spring, Summer, Fall")

	StardewCrops.AddCommand(SearchCmd)
}
