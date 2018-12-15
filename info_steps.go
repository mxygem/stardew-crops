package main

import (
	"strings"

	"github.com/jaysonesmith/stardew-crops/utils"
)

func (sc *ScenarioContext) InfoCrop(crop string) error {
	sc.STDOut = InfoCommand(crop)

	return nil
}

func (sc *ScenarioContext) InfoNoArgs() error {
	sc.STDOut = InfoCommand("")

	return nil
}

func (sc *ScenarioContext) MatchCropInfo(crop string) error {
	expected := utils.Open("./test_data/crop_garlic_raw.json")
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchNoCropSpecifiedMessage() error {
	expected := `"No crop name specified, please try again"`
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchInfoCropNotFound() error {
	expected := `"Unable to find matching crop for banana"`
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}
