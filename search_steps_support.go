package main

import (
	"github.com/jaysonesmith/stardew-crops/commands"
	"github.com/jaysonesmith/stardew-crops/utils"
	"github.com/spf13/cobra"
)

// SearchCommand wraps the call to the Search processor
// to capture the stdout
func SearchCommand(args ...string) string {
	so, r, w := utils.STDOutUp()

	commands.Search(&cobra.Command{}, args)

	return string(utils.STDOutDown(so, r, w))
}
