package steps

import (
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/jaysonesmith/stardew-crops/commands"
)

func ThePrettyFormatterIsDesired() error {
	return godog.ErrPending
}

func InfoForGarlicIsRequested() error {
	commands.Execute()

	return godog.ErrPending
}

func TheOutputMustMatch(arg1 *gherkin.DocString) error {
	return godog.ErrPending
}
