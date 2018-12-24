package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/tidwall/pretty"
)

func Format(data interface{}, f string) *bytes.Buffer {
	var b []byte

	b, err := json.Marshal(data)
	if err != nil {
		b = []byte(fmt.Sprintf("unable to marshal output: %s", err.Error()))
	}

	switch f {
	case "pretty":
		fmt.Println("pretty format requested")
		fmt.Println("TODO: add pretty formatter")
	case "json":
		b = pretty.Pretty(b)
	case "raw", "":
		b = pretty.Ugly(b)
	}

	return bytes.NewBuffer(b)
}

// Print provides a single source of completion for the CLI
func Print(o *bytes.Buffer) {
	o.WriteTo(os.Stdout)
}
