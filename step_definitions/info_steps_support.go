package stepdefinitions

// InfoCommand wraps the call to the Info processor
// to capture the stdout
func (sc *ScenarioContext) InfoCommand(args ...string) {
	sc.STDOut = StardewCropsCommand(args)
}
