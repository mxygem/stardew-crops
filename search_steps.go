package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jaysonesmith/stardew-crops/utils"
)

func (sc *ScenarioContext) SearchNoArgs() error {
	StardewCropsCommand([]string{"search", "--bundle"})

	return nil
}

func (sc *ScenarioContext) Search(flag, value string) error {
	args := []string{"search", fmt.Sprintf(`--%s=%s`, flag, value)}
	sc.STDOut = StardewCropsCommand(args)

	return nil
}

func (sc *ScenarioContext) GrowthSearch(timeType string, value int64) error {
	flag := map[string]string{"greater": "growthgt", "less": "growthlt"}

	return sc.Search(flag[timeType], strconv.Itoa(int(value)))
}

func (sc *ScenarioContext) MatchBundleCrops(bundle string) error {
	expected := utils.Open("./test_data/search/bundle/summerCrops.json")
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchNotFound() error {
	expected := `"No matching crops found"`
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchGrowthResults(timeType string) error {
	fileName := map[string]string{"less": "lessThanFive", "more": "greaterThanFive"}
	expected := utils.Open(fmt.Sprintf(`./test_data/search/growth/%s.json`, fileName[timeType]))
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchSeasonResults(season string) error {
	expected := utils.Open(fmt.Sprintf(`./test_data/search/season/%s.json`, season))
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

// remove below?

// func (sc *ScenarioContext) MatchAllCropsBySeason() error {
// 	expected := utils.Open("./test_data/search/season/all.json")
// 	actual := strings.TrimSpace(sc.STDOut)

// 	return utils.AssertMatch(expected, actual)
// }

// func (sc *ScenarioContext) MatchSeasonCrops(season string) error {
// 	expected := utils.Open(fmt.Sprintf(`./test_data/search/season/%s.json`, season))
// 	actual := strings.TrimSpace(sc.STDOut)

// 	return utils.AssertMatch(expected, actual)
// }

// func (sc *ScenarioContext) MatchUnknownCropMessage() error {
// 	expected := `"Unknown season for breakfast"`
// 	actual := strings.TrimSpace(sc.STDOut)

// 	return utils.AssertMatch(expected, actual)
// }
