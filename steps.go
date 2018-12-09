package main

import (
	"fmt"

	"github.com/DATA-DOG/godog/gherkin"
	"github.com/jaysonesmith/stardew-crops/processors"
)

func InfoCommand(crop string) error {
	processors.Info()

	return nil
}

func TheOutputMustMatch(expected *gherkin.DocString) error {
	actual := fmt.Sprintf("%s\n", &cliOut)
	if expected.Content == actual {
		return nil
	}

	return fmt.Errorf("expected content and found content do not match!\n expected: %s\n found: %s", expected.Content, actual)
}
