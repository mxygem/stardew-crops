package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

// Print provides a single source of completion for the CLI
func Print(data interface{}, f string) {
	var b []byte
	var err error
	var out bytes.Buffer

	b, err = json.Marshal(data)
	if err != nil {
		// Add error to buffer and write
		fmt.Errorf("unable to marshal output: %s", err.Error())
	}

	switch f {
	case "pretty":
		fmt.Println("pretty format requested")
		fmt.Println("TODO: add pretty formatter")
	case "json":
		out, err = indentedJSON(b)
	case "raw", "":
		b, err = json.Marshal(data)
		if err != nil {
			fmt.Println("unable to process output, things are borked")
		}

		if err := json.Compact(&out, b); err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("unknown format requested")
		fmt.Println("TODO: point default to pretty")
	}

	if err != nil {
		// Add error to buffer and write
		fmt.Errorf("unable to process output: %s", err.Error())
	}

	out.WriteTo(os.Stdout)
}

func indentedJSON(data []byte) (bytes.Buffer, error) {
	var out bytes.Buffer

	err := json.Indent(&out, data, "", "\t")
	if err != nil {
		return out, fmt.Errorf("unable to indent output: %s", err.Error())
	}

	return out, nil
}
