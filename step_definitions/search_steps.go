package stepdefinitions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jaysonesmith/stardew-crops/utils"
)

func (sc *ScenarioContext) Search(flag, value string) error {
	args := []string{"search", fmt.Sprintf(`--%s=%s`, flag, value)}
	sc.STDOut = StardewCropsCommand(args)

	return nil
}

func (sc *ScenarioContext) DoubleSearch(flag1, value1, flag2, value2 string) error {
	args := []string{"search", fmt.Sprintf(`--%s=%s`, flag1, value1), fmt.Sprintf(`--%s=%s`, flag2, value2)}
	sc.STDOut = StardewCropsCommand(args)

	return nil
}

func (sc *ScenarioContext) TripleSearch(flag1, value1, flag2, value2, flag3, value3 string) error {
	args := []string{
		"search",
		fmt.Sprintf(`--%s=%s`, flag1, value1),
		fmt.Sprintf(`--%s=%s`, flag2, value2),
		fmt.Sprintf(`--%s=%s`, flag3, value3),
	}
	sc.STDOut = StardewCropsCommand(args)

	return nil
}

func (sc *ScenarioContext) GrowthSearch(timeType string, value int64) error {
	flag := map[string]string{"greater": "growthgt", "less": "growthlt"}

	return sc.Search(flag[timeType], strconv.Itoa(int(value)))
}

func (sc *ScenarioContext) TrellisSearch(t string) error {
	flag := map[string]string{"grow": "true", "do not grow": "false"}
	return sc.Search("trellis", flag[t])
}

func (sc *ScenarioContext) ContinuousSearch(t string) error {
	flag := map[string]string{"continuously harvestable": "true", "single harvest": "false"}
	return sc.Search("continuous", flag[t])
}

func (sc *ScenarioContext) MatchBundleCrops(bundle string) error {
	expected := utils.Open(".././test_data/search/bundle/summerCrops.json")
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchNotFound() error {
	expected := `"No matching crops found"`
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchGrowthResults(g int64) error {
	expected := utils.Open(fmt.Sprintf(`.././test_data/search/growth/%v_days.json`, g))
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchGrowthGTLTResults(timeType string) error {
	fileName := map[string]string{"less": "lessThanFive", "more": "greaterThanFive"}
	expected := utils.Open(fmt.Sprintf(`.././test_data/search/growth/%s.json`, fileName[timeType]))
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchSeasonResults(season string) error {
	expected := utils.Open(fmt.Sprintf(`.././test_data/search/season/%s.json`, season))
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchTrellisResults(presence string) error {
	fileName := map[string]string{"do not grow": "off_trellis", "grow": "on_trellis"}
	expected := utils.Open(fmt.Sprintf(`.././test_data/search/trellis/%s.json`, fileName[presence]))
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchContinuousResults(continuous string) error {
	fileName := map[string]string{"are": "continuous", "are not": "single"}
	expected := utils.Open(fmt.Sprintf(`.././test_data/search/continuous/%s.json`, fileName[continuous]))
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchSingleCrop(c string) error {
	expected := utils.Open(fmt.Sprintf(`.././test_data/search/crops/%v.json`, c))
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}
