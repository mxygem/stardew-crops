package main

import (
	"github.com/jaysonesmith/stardew-crops/commands"
	"github.com/jaysonesmith/stardew-crops/utils"
)

// StardewCropsCommand wraps the call to the root
// command and returns the captured output
func StardewCropsCommand(args []string) string {
	so, r, w := utils.STDOutUp()

	cmd := commands.StardewCrops
	cmd.SetArgs(args)
	cmd.Execute()

	return string(utils.STDOutDown(so, r, w))
}
