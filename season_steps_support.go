package main

import (
	"fmt"

	"github.com/jaysonesmith/stardew-crops/commands"
	"github.com/jaysonesmith/stardew-crops/utils"
	"github.com/spf13/cobra"
)

// SeasonCommand wraps the call to the Info processor
// to capture the stdout
func SeasonCommand(args ...string) string {
	so, r, w := utils.STDOutUp()

	commands.Season(&cobra.Command{}, args)

	return string(utils.STDOutDown(so, r, w))
}

func SeasonTestData(season string) string {
	return utils.Open(fmt.Sprintf(`./test_data/%s`, season))
}

func AssertMatch(expected, actual string) error {
	if expected == actual {
		return nil
	}

	return fmt.Errorf("expected content and found content do not match!\n expected: %s\n found: %s", expected, actual)
}
