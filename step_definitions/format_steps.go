package stepdefinitions

import (
	"fmt"
	"strings"

	"github.com/DATA-DOG/godog/gherkin"
	"github.com/jaysonesmith/stardew-crops/utils"
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

func (sc *ScenarioContext) MatchOutput(e *gherkin.DocString) error {
	expected := strings.TrimSpace(e.Content)
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}
