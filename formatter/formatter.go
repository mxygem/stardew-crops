package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
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

func prettyFormat(d data.CropData) []byte {
	funcs := template.FuncMap{
		"title":     title,
		"season":    season,
		"add":       add,
		"safe":      safe,
		"lineSplit": lineSplit,
		"pad":       pad,
		"derefTCP":  derefTravelingCartPrices,
		"listData":  listData,
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

// lineBreaks manages wrapping an input string
// based on the input line length while also
// preventing words from being split in half
// by splitting at the closest previous space
func lineBreaks(s string, l int) []string {
	sb := []byte(s)
	sl := len(sb)

	if l >= sl {
		return []string{s}
	}

	out := []string{}
	start := 0
	end := l
	for end <= sl {
		// padding end number if the current string
		// starts with a space
		if sb[start] == 32 {
			end++
		}

		// check if we're trying to split inside a
		// word and update the end of our cut with
		// the position of the last space
		if end < sl {
			if sb[end] != 32 && sb[end-1] != 32 {
				end = lastSpace(sb, end)
			}
		}

		// make the cut
		var sub []byte
		if end <= sl {
			sub = sb[start:end]
		}

		// don't append strings
		if string(sub) != "" {
			out = append(out, strings.TrimSpace(string(sub)))
		}

		// update for the next line
		start = end
		end += l

		// handle tracking for a cut that ends in a
		// space
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

func derefTravelingCartPrices(i *data.Prices) data.TravelingCartPrices {
	if i.TravelingCart != nil {
		return *i.TravelingCart
	}
	return data.TravelingCartPrices{}
}

func pad(l int, s string) string {
	return fmt.Sprintf("%s%s", s, bytes.Repeat([]byte(" "), (l-len([]byte(s)))))
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

// ListData is a collection of data for use in
// the pretty template
type ListData struct {
	Name  string
	Items []string
}

func listData(name string, item []string) ListData {
	return ListData{
		Name:  name,
		Items: item,
	}
}
