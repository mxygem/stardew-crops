package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/tidwall/pretty"
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
		out = *bytes.NewBuffer(pretty.Pretty(b))
	case "raw", "":
		out = *bytes.NewBuffer(pretty.Ugly(b))
	}

	if err != nil {
		// Add error to buffer and write
		fmt.Errorf("unable to process output: %s", err.Error())
		os.Exit(1)
	}

	out.WriteTo(os.Stdout)
	os.Exit(0)
}
