package main

import (
	"github.com/DATA-DOG/godog"
)

func FeatureContext(s *godog.Suite) {

	sc := ScenarioContext{}

	s.AfterScenario(func(interface{}, error) {
		sc = ScenarioContext{}
	})

	s.Step(`^info for (\w+) is requested$`, sc.InfoCommand)
	s.Step(`^the output must match$`, sc.MatchDocString)
	s.Step(`^the output must match the file "([^"]*)"$`, sc.MatchFile)
}
