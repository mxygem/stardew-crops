package main

import (
	"github.com/jaysonesmith/stardew-crops/processors"
	"github.com/jaysonesmith/stardew-crops/utils"
)

// SeasonCommand wraps the call to the Info processor
// to capture the stdout
func SeasonCommand(args ...string) string {
	so, r, w := utils.STDOutUp()

	processors.Season(args...)

	return string(utils.STDOutDown(so, r, w))
}
