package main

func (sc *ScenarioContext) InfoCrop(crop string) error {
	sc.STDOut = InfoCommand(crop)

	return nil
}

func (sc *ScenarioContext) InfoNoArgs() error {
	sc.STDOut = InfoCommand("")

	return nil
}
