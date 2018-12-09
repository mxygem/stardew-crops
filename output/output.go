package output

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Print provides a single source of completion for the CLI
func Print(data interface{}) {
	out, err := json.Marshal(data)
	if err != nil {
		fmt.Println("unable to process output, things are borked")
	}

	dst := new(bytes.Buffer)
	if err := json.Compact(dst, out); err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(out))
}
