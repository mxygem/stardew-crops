package output

import (
	"bytes"
	"os"
)

// Print provides a single source of completion
// for the CLI
func Print(o *bytes.Buffer) {
	o.WriteTo(os.Stdout)
}
