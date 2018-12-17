package main

import (
	"log"
	"os/exec"
)

// StardewCropsCommand wraps the call to the root
// command and returns the captured output
func StardewCropsCommand(args []string) string {
	// so, r, w := utils.STDOutUp()

	// cmd := commands.StardewCrops
	// cmd.SetArgs(args)
	// cmd.Execute()

	// return string(utils.STDOutDown(so, r, w))

	cmd := exec.Command("stardew-crops", args[0], args[1])
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	return string(out)
}
