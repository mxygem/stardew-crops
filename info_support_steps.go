package main

import (
	"github.com/jaysonesmith/stardew-crops/commands"
	"github.com/jaysonesmith/stardew-crops/utils"
	"github.com/spf13/cobra"
)

// InfoCommand wraps the call to the Info processor
// to capture the stdout
func InfoCommand(args ...string) string {
	so, r, w := utils.STDOutUp()

	commands.Info(&cobra.Command{}, args)

	return string(utils.STDOutDown(so, r, w))
}
