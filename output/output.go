package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"os"

	"github.com/jaysonesmith/stardew-crops/utils"
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
		b = prettyFormat(b)
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

func prettyFormat(b []byte) []byte {
	funcs := template.FuncMap{
		// "safe": safe,
	}

	t := template.New("pretty").Funcs(funcs)
	template.Must(t.Parse(utils.Open("../template/pretty.tmpl")))

	out := new(bytes.Buffer)
	err := t.Execute(out, b)
	if err != nil {
		panic(err)
	}

	return out.Bytes()
}
