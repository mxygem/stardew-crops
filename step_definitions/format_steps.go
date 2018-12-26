package stepdefinitions

import (
	"fmt"
	"strings"

	"github.com/jaysonesmith/stardew-crops/utils"
	"github.com/tidwall/pretty"
)

func (sc *ScenarioContext) FormatRequester(key, value, formatter string) error {
	args := []string{key, value, formatter}
	sc.STDOut = StardewCropsCommand(args)

	return nil
}

func (sc *ScenarioContext) SearchFormat(format string) error {
	args := []string{"search", fmt.Sprintf("--format=%s", format), "--continuous=true", "--growth=13"}
	sc.STDOut = StardewCropsCommand(args)

	return nil
}

func (sc *ScenarioContext) CheckPretty() error {
	b := []byte(sc.STDOut)
	u := pretty.Ugly(b)

	if string(b) == string(u) {
		return fmt.Errorf("output doesn't appear to be pretty json: %s", sc.STDOut)
	}

	return nil
}

func (sc *ScenarioContext) CheckRaw() error {
	u := pretty.Ugly([]byte(sc.STDOut))

	if sc.STDOut != string(u) {
		return fmt.Errorf("output doesn't appear to be raw json: %s", sc.STDOut)
	}

	return nil
}

func (sc *ScenarioContext) InfoFormat(c, f string) error {
	args := []string{"info", fmt.Sprintf("--format=%s", f), c}
	sc.STDOut = StardewCropsCommand(args)

	return nil
}

func (sc *ScenarioContext) MatchInfoFormat(c, f string) error {
	expected := strings.TrimSpace(utils.Open(fmt.Sprintf(".././test_data/format/%s_info_%s.txt", f, c)))
	actual := strings.TrimSpace(sc.STDOut)

	if expected != actual {
		return fmt.Errorf("actual output does not match expected.\nexpected:%s\nactual:%s", expected, actual)
	}

	return nil
}
