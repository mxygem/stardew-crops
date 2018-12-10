package main

import (
	"github.com/jaysonesmith/stardew-crops/processors"
	"github.com/jaysonesmith/stardew-crops/utils"
)

// InfoCommand wraps the call to the Info processor
// to capture the stdout
func InfoCommand(args ...string) string {
	so, r, w := utils.STDOutUp()

	processors.Info(args...)

	return string(utils.STDOutDown(so, r, w))
}
