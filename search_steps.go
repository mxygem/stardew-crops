package main

import (
	"fmt"
	"strings"

	"github.com/jaysonesmith/stardew-crops/utils"
)

func (sc *ScenarioContext) SearchNoArgs() error {
	sc.STDOut = SearchCommand("")

	return nil
}

func (sc *ScenarioContext) Search(flag, value string) error {
	sc.STDOut = SearchCommand(fmt.Sprintf("--%s %s", flag, value))

	return nil
}

func (sc *ScenarioContext) MatchBundleCrops(bundle string) error {
	expected := `["blueberry", "hot pepper"]`
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchNotFound() error {
	expected := `"Unable to find matching crop for Pasta"`
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchNoValueSpecifiedMessage() error {
	expected := `"A value must be provided for each parameter specified"`
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

// remove below?

func (sc *ScenarioContext) MatchAllCropsBySeason() error {
	expected := utils.Open("./test_data/seasons/all.json")
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchSeasonCrops(season string) error {
	expected := utils.Open(fmt.Sprintf(`./test_data/seasons/%s.json`, season))
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchUnknownCropMessage() error {
	expected := `"Unknown season for breakfast"`
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}
