package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/jaysonesmith/stardew-crops/steps"
)

func FeatureContext(s *godog.Suite) {
	s.Step(`^the pretty formatter is desired$`, steps.ThePrettyFormatterIsDesired)
	s.Step(`^info for garlic is requested$`, steps.InfoForGarlicIsRequested)
	s.Step(`^the output must match$`, steps.TheOutputMustMatch)
}
