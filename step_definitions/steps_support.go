package stepdefinitions

import (
	"fmt"
	"os"
	"os/exec"
)

// StardewCropsCommand wraps the call to the root
// command and returns the captured output
func StardewCropsCommand(args []string) string {
	cmd := exec.Command("stardew-crops", args[0], args[1])
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
		os.Exit(1)
	}

	return string(out)
}
