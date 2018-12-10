package main

func (sc *ScenarioContext) InfoCrop(crop string) error {
	sc.STDOut = InfoCommand(crop)

	return nil
}
