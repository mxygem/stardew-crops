package main

import (
	"fmt"
	"strings"

	"github.com/DATA-DOG/godog/gherkin"
	"github.com/jaysonesmith/stardew-crops/utils"
)

func AssertMatch(expected, actual string) error {
	if expected == actual {
		return nil
	}

	return fmt.Errorf("expected content and found content do not match!\n expected: %s\n found: %s", expected, actual)
}

func (sc *ScenarioContext) MatchString(e string) error {
	actual := strings.TrimSpace(sc.STDOut)
	expected := strings.TrimSpace(e)
	if expected == actual {
		return nil
	}

	return fmt.Errorf("expected content and found content do not match!\n expected: %s\n found: %s", expected, actual)
}

func (sc *ScenarioContext) MatchDocString(e *gherkin.DocString) error {
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

	fmt.Println("expect", expected)
	fmt.Println("actual", actual)
	return fmt.Errorf("expected content and found content do not match\n%v", utils.Diff(expected, actual))
}
