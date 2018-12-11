package main

import (
	"github.com/DATA-DOG/godog"
)

func FeatureContext(s *godog.Suite) {

	sc := ScenarioContext{}

	s.AfterScenario(func(interface{}, error) {
		sc = ScenarioContext{}
	})

	s.Step(`^info for (\w+) is requested$`, sc.InfoCrop)
	s.Step(`^info is requested$`, sc.InfoNoArgs)
	s.Step(`^the output must match ("[^"]*")$`, sc.MatchString)
	s.Step(`^the output must match$`, sc.MatchDocString)
	s.Step(`^the output must match the file "([^"]*)"$`, sc.MatchFile)
	s.Step(`^the (\w+) season is requested$`, sc.Season)
	s.Step(`^season is requested$`, sc.SeasonNoArgs)
}
