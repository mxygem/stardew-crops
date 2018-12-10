package main

import (
	"fmt"
	"strings"

	"github.com/DATA-DOG/godog/gherkin"
	"github.com/jaysonesmith/stardew-crops/utils"
)

func (sc *ScenarioContext) MatchSTDOut(e *gherkin.DocString) error {
	actual := strings.TrimSpace(sc.STDOut)
	expected := strings.TrimSpace(e.Content)
	if expected == actual {
		return nil
	}

	return fmt.Errorf("expected content and found content do not match!\n expected: %s\n found: %s", expected, actual)
}

func (sc *ScenarioContext) MatchFile(path string) error {
	actual := strings.TrimSpace(sc.STDOut)
	expected := utils.Open(path)

	if strings.Compare(expected, actual) == 0 {
		return nil
	}

	return fmt.Errorf("expected content and found content do not match\n%v", utils.Diff(expected, actual))
}
