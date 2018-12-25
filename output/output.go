package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/jaysonesmith/stardew-crops/data"
	"github.com/jaysonesmith/stardew-crops/utils"
	"github.com/tidwall/pretty"
)

func Format(d interface{}, f string) *bytes.Buffer {
	var b []byte

	b, err := json.Marshal(d)
	if err != nil {
		b = []byte(fmt.Sprintf("unable to marshal output: %s", err.Error()))
	}

	switch f {
	case "pretty":
		b = prettyFormat(d.(data.CropData))
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

func prettyFormat(d data.CropData) []byte {
	funcs := template.FuncMap{
		"title":     title,
		"season":    season,
		"add":       add,
		"safe":      safe,
		"noteSplit": NoteSplit,
	}

	t := template.New("pretty").Funcs(funcs)
	template.Must(t.Parse(utils.Open("../template/pretty.tmpl")))

	out := new(bytes.Buffer)
	err := t.Execute(out, d)
	if err != nil {
		panic(err)
	}

	return out.Bytes()
}

func title(s string) string {
	return strings.Title(s)
}

func season(s []string) string {
	return title(strings.Join(s, ", "))
}

func add(n1, n2 int) int {
	return n1 + n2
}

func safe(s string) template.HTML {
	return template.HTML(s)
}

func NoteSplit(s string) string {
	sb := []byte(s)
	l := len(sb)
	var o strings.Builder
	fmt.Println("length:", l)
	if l == 42 {
		fmt.Fprintf(&o, "║   * %s ║", s)
	} else if l < 42 {
		out := pad(42, s)
		fmt.Fprintf(&o, "║   * %s ║", out)
	} else if l > 42 {
		fmt.Fprintf(&o, "║   * %s ║", sb[0:42])
		fmt.Fprint(&o, "\n")
		fmt.Fprintf(&o, "║     %s ║", pad(42, string(sb[43:81])))
		fmt.Fprint(&o, "\n")
		fmt.Fprintf(&o, "║     %s ║", pad(42, string(sb[82:87])))
	}

	return o.String()
}

func pad(l int, s string) string {
	return fmt.Sprintf("%s%s", s, bytes.Repeat([]byte(" "), (l-len([]byte(s)))))
}
