package main

import (
	"github.com/jaysonesmith/stardew-crops/commands"
	"github.com/jaysonesmith/stardew-crops/utils"
	"github.com/spf13/cobra"
)

// SeasonCommand wraps the call to the Info processor
// to capture the stdout
func SeasonCommand(args ...string) string {
	so, r, w := utils.STDOutUp()

	commands.Season(&cobra.Command{}, args)

	return string(utils.STDOutDown(so, r, w))
}
