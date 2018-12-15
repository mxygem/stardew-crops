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
func userSetFlags(cmd *cobra.Command) map[string]interface{} {
	userFlags := make(map[string]interface{}, 0)
	cmd.Flags().Visit(func(f *pflag.Flag) {
		if f.Changed {
			userFlags[f.Name] = f.Value
		}
	})

	return userFlags
}

func init() {
	// set flags
	searchCmd.Flags().StringP("bundle", "b", "", "Specify a bundle name to search for")
	searchCmd.Flags().BoolP("continuous", "c", false, "Search for crops that grow continuously")
	searchCmd.Flags().Int64P("growthgt", "g", 0, "Search for crops that take the specified time or longer to grow")
	searchCmd.Flags().Int64P("growthlt", "l", 0, "Search for crops that take the specified time or less to grow")
	searchCmd.Flags().StringP("season", "s", "", "Search by a season. Options: Spring, Summer, Fall")
	searchCmd.Flags().BoolP("trellis", "t", false, "Search for crops. Options: Spring, Summer, Fall")

	stardewCrops.AddCommand(searchCmd)
}
