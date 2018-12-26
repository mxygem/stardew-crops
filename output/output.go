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
		"lineSplit": lineSplit,
		"pad":       pad,
		"derefTCP":  derefTravelingCartPrices,
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

func derefTravelingCartPrices(i *data.Prices) data.TravelingCartPrices {
	if i.TravelingCart != nil {
		return *i.TravelingCart
	}
	return data.TravelingCartPrices{}
}

func lineSplit(s string, l int) []string {
	sl := len([]byte(s))
	var sb []string

	if sl == l {
		sb = append(sb, fmt.Sprintf("   * %s ", s))
	} else if sl < l {
		sb = append(sb, fmt.Sprintf("   * %s ", pad(l, s)))
	} else if sl > l {
		lines := lineBreaks(s, l)
		for i, line := range lines {
			if i == 0 {
				sb = append(sb, fmt.Sprintf("   * %s ", pad(l, line)))
			} else {
				sb = append(sb, fmt.Sprintf("     %s ", pad(l, line)))
			}
		}
	}

	return sb
}

func pad(l int, s string) string {
	return fmt.Sprintf("%s%s", s, bytes.Repeat([]byte(" "), (l-len([]byte(s)))))
}

func lineBreaks(s string, l int) []string {
	sb := []byte(s)
	sl := len(sb)
	// if the line length is greater than the
	// input, nothing needs to happen
	if l >= sl {
		return []string{s}
	}

	out := []string{}
	start := 0
	end := l
	for end <= sl {
		// pad ending if starting with a space
		if sb[start] == 32 {
			end++
		}

		// make sure we're not cutting in the middle
		// of a word and instead find the last space
		// before the desired cut
		if end < sl {
			if sb[end] != 32 && sb[end-1] != 32 {
				end = lastSpace(sb, end)
			}
		}

		// grab the content
		var sub []byte
		if end <= sl {
			sub = sb[start:end]
		}

		if string(sub) != "" {
			out = append(out, strings.TrimSpace(string(sub)))
		}

		// update for the next line
		start = end
		end += l
		if len(sub) > 0 {
			if sub[(len(sub)-1)] == 32 {
				end++
			}
		}
	}

	// grab leftover content
	if end > sl && start < sl {
		out = append(out, strings.TrimSpace(string(sb[start:])))
	}

	return out
}

func lastSpace(sb []byte, end int) int {
	for i := end; i > 0; i-- {
		if sb[i] == 32 {
			return i
		}
	}
	return 0
}
