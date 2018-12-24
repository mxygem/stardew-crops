package stepdefinitions

import (
	"strings"

	"github.com/jaysonesmith/stardew-crops/utils"
)

func (sc *ScenarioContext) InfoCrop(crop string) error {
	args := []string{"info", crop}
	sc.STDOut = StardewCropsCommand(args)

	return nil
}

func (sc *ScenarioContext) InfoNoArgs() error {
	args := []string{"info"}
	sc.STDOut = StardewCropsCommand(args)

	return nil
}

func (sc *ScenarioContext) MatchCropInfo(crop string) error {
	expected := strings.TrimSpace(utils.Open(".././test_data/format/garlic_raw.json"))
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
