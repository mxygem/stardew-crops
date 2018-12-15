package main

import (
	"github.com/DATA-DOG/godog"
)

func FeatureContext(s *godog.Suite) {

	sc := ScenarioContext{}

	s.AfterScenario(func(interface{}, error) {
		sc = ScenarioContext{}
	})

	// Info steps
	s.Step(`^info is requested$`, sc.InfoNoArgs)
	s.Step(`^info for (\w+) is requested$`, sc.InfoCrop)
	s.Step(`^a message informing no crop was specified must be returned$`, sc.MatchNoCropSpecifiedMessage)
	s.Step(`^the info for (\w+) must be returned$`, sc.MatchCropInfo)
	s.Step(`^a message informing no matching crop was found must be returned$`, sc.MatchInfoCropNotFound)

	// Season steps
	s.Step(`^the (\w+) season\'s crops (?:is|are) requested$`, sc.Season)
	s.Step(`^season is requested without a specified season$`, sc.SeasonNoArgs)
	s.Step(`^all crops grouped by season must be returned$`, sc.MatchAllCropsBySeason)
	s.Step(`^only the (\w+) crops must be returned$`, sc.MatchSeasonCrops)
	s.Step(`^a message informing no matched season was found must be returned$`, sc.MatchUnknownCropMessage)
}
