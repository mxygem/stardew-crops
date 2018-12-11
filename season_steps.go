package main

import (
	"fmt"
	"strings"

	"github.com/jaysonesmith/stardew-crops/utils"
)

func (sc *ScenarioContext) Season(crop string) error {
	sc.STDOut = SeasonCommand(crop)

	return nil
}

func (sc *ScenarioContext) SeasonNoArgs() error {
	sc.STDOut = SeasonCommand("")

	return nil
}

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
