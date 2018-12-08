package output

import (
	"encoding/json"
	"fmt"
)

// Print provides a single source of completion for the CLI
func Print(data interface{}) {
	out, err := json.Marshal(data)
	if err != nil {
		fmt.Println("unable to process output, things are borked")
	}

	fmt.Println(string(out))
}
