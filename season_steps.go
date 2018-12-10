package main

func (sc *ScenarioContext) Season(crop string) error {
	sc.STDOut = SeasonCommand(crop)

	return nil
}

func (sc *ScenarioContext) SeasonNoArgs() error {
	sc.STDOut = SeasonCommand("")

	return nil
}
