package main

import (
	"fmt"
	"strings"
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
	actual := strings.TrimSpace(sc.STDOut)
	expected := `{"spring":["garlic","potato"],"summer":["blueberry","radish","starfruit"],"fall":["cranberries","yams"]}`

	if expected == actual {
		return nil
	}

	return fmt.Errorf("expected content and found content do not match!\n expected: %s\n found: %s", expected, actual)
}

func (sc *ScenarioContext) MatchSeasonCrops(season string) error {
	actual := strings.TrimSpace(sc.STDOut)
	expected := `{"summer":["blueberry","radish","starfruit"]}`

	if expected == actual {
		return nil
	}

	return fmt.Errorf("expected content and found content do not match!\n expected: %s\n found: %s", expected, actual)
}

func (sc *ScenarioContext) MatchUnknownCropMessage() error {
	actual := strings.TrimSpace(sc.STDOut)
	expected := `"Unknown season for breakfast"`

	if expected == actual {
		return nil
	}

	return fmt.Errorf("expected content and found content do not match!\n expected: %s\n found: %s", expected, actual)
}
