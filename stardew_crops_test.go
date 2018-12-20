package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/jaysonesmith/stardew-crops/step_definitions"
)

func FeatureContext(s *godog.Suite) {
	var sc stepdefinitions.ScenarioContext

	s.AfterScenario(func(interface{}, error) {
		sc = stepdefinitions.ScenarioContext{STDOut: ""}
	})

	// Info steps
	s.Step(`^info is requested$`, sc.InfoNoArgs)
	s.Step(`^info for (\w+) is requested$`, sc.InfoCrop)

	s.Step(`^the info for (\w+) must be returned$`, sc.MatchCropInfo)
	s.Step(`^a message informing no crop was specified must be returned$`, sc.MatchNoCropSpecifiedMessage)
	s.Step(`^a message informing no matching crop was found must be returned$`, sc.MatchInfoCropNotFound)

	// Search steps
	s.Step(`^a search by (\w+) for ([^"]*) is performed$`, sc.Search)
	s.Step(`^a search by (\w+) (\w+) is performed$`, sc.Search)
	s.Step(`^a search by (\w+) than growth time of (\d+) days is performed$`, sc.GrowthSearch)
	s.Step(`^a search for crops that (do not grow|grow) on a trellis is performed$`, sc.TrellisSearch)
	s.Step(`^a search for (continuously harvestable|single harvest) crops is performed$`, sc.ContinuousSearch)

	s.Step(`^a list of the crops available in the ([^"]*) bundle must be returned$`, sc.MatchBundleCrops)
	s.Step(`^an error indicating that no match was found must be returned$`, sc.MatchNotFound)
	s.Step(`^a list of crops that grow in 5 days or (\w+) must be returned$`, sc.MatchGrowthResults)
	s.Step(`^a list of crops that grow in (\w+) must be returned$`, sc.MatchSeasonResults)
	s.Step(`^a list of crops that (do not grow|grow) on a trellis are returned$`, sc.MatchTrellisResults)
	s.Step(`^a list of crops that (are|are not) continuously harvestable are returned$`, sc.MatchContinuousResults)
}
