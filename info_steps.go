package main

import (
	"github.com/jaysonesmith/stardew-crops/processors"
	"github.com/jaysonesmith/stardew-crops/utils"
)

func (sc *ScenarioContext) InfoCommand(crop string) error {
	so, r, w := utils.STDOutUp()
	processors.Info()
	sc.STDOut = string(utils.STDOutDown(so, r, w))

	return nil
}
