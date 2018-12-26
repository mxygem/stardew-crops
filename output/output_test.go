package output_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/jaysonesmith/stardew-crops/data"
	"github.com/jaysonesmith/stardew-crops/output"
	"github.com/jaysonesmith/stardew-crops/utils"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/pretty"
)

func TestFormat(t *testing.T) {
	testCases := []struct {
		name     string
		data     interface{}
		f        string
		expected *bytes.Buffer
	}{
		{
			name:     "Empty",
			data:     "",
			f:        "",
			expected: bytes.NewBuffer([]byte(`""`)),
		},
		{
			name: "Raw - Info response CropData Starfruit",
			data: data.Crop{
				Name: "garlic",
				Info: data.Info{
					Description: "Adds a wonderful zestiness to dishes. High quality garlic can be pretty spicy.",
					Seed:        "garlic seeds",
					GrowthTime:  4,
					Season:      []string{"spring"},
				},
				SeedPrices: data.Prices{GeneralStore: 40},
				Recipes:    []string{"Escargot", "Fiddlehead Risotto", "Oil of Garlic"},
				Notes:      []string{"Only available starting in year 2"},
			},
			f:        "raw",
			expected: bytes.NewBuffer([]byte(`{"name":"garlic","info":{"description":"Adds a wonderful zestiness to dishes. High quality garlic can be pretty spicy.","seed":"garlic seeds","growth_time":4,"season":["spring"]},"seed_prices":{"general_store":40},"recipes":["Escargot","Fiddlehead Risotto","Oil of Garlic"],"notes":["Only available starting in year 2"]}`)),
		},
		{
			name: "JSON - Info response CropData Starfruit",
			data: data.Crop{
				Name: "garlic",
				Info: data.Info{
					Description: "Adds a wonderful zestiness to dishes. High quality garlic can be pretty spicy.",
					Seed:        "garlic seeds",
					GrowthTime:  4,
					Season:      []string{"spring"},
				},
				SeedPrices: data.Prices{GeneralStore: 40},
				Recipes:    []string{"Escargot", "Fiddlehead Risotto", "Oil of Garlic"},
				Notes:      []string{"Only available starting in year 2"},
			},
			f:        "json",
			expected: bytes.NewBuffer(pretty.Pretty([]byte(`{"name":"garlic","info":{"description":"Adds a wonderful zestiness to dishes. High quality garlic can be pretty spicy.","seed":"garlic seeds","growth_time":4,"season":["spring"]},"seed_prices":{"general_store":40},"recipes":["Escargot","Fiddlehead Risotto","Oil of Garlic"],"notes":["Only available starting in year 2"]}`))),
		},
		{
			name: "Pretty - Info response Starfruit",
			data: data.CropData{
				Crops: []data.Crop{data.Crop{
					Name: "starfruit",
					Info: data.Info{
						Description: "An extremely juicy fruit that grows in hot, humid weather. Slightly sweet with a sour undertone.",
						Seed:        "starfruit seeds",
						GrowthTime:  13,
						Season:      []string{"summer"},
						Continual:   false,
					},
					SeedPrices: data.Prices{Oasis: 400},
					Quests:     []string{"A Soldier's Star"},
					Notes:      []string{"Starfruit produces Artisan Goods that have some of the highest sell values in the game.", "Starfruit is required to build a Junimo Hut, purchased from the Wizard's Tower."},
				}}},
			f:        "pretty",
			expected: bytes.NewBuffer([]byte(strings.TrimSpace(utils.Open(".././test_data/format/pretty_info_starfruit.txt")))),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := output.Format(tc.data, tc.f)

			assert.Equal(t, tc.expected.String(), actual.String())
		})
	}
}

func TestLineSplit(t *testing.T) {
	testCases := []struct {
		name       string
		input      string
		expected   []string
		lineLength int
	}{
		{
			name:       "Note exactly 42 characters",
			input:      "Starfruit produces Artisan Goods that have",
			lineLength: 42,
			expected:   []string{"   * Starfruit produces Artisan Goods that have "},
		},
		{
			name:       "Note under 42 characters to force padding",
			input:      "Starfruit produces Artisan",
			lineLength: 42,
			expected:   []string{"   * Starfruit produces Artisan                 "},
		},
		{
			name:       "Multiline and padding of note with 87 characters",
			input:      "Starfruit produces Artisan Goods that have some of the highest sell values in the game.",
			lineLength: 42,
			expected: []string{"   * Starfruit produces Artisan Goods that have ",
				"     some of the highest sell values in the     ",
				"     game.                                      ",
			},
		},
		{
			name:       "Multiline and padding of note with 60 characters",
			input:      "Starfruit produces Artisan Goods that have some of the highe",
			lineLength: 42,
			expected: []string{
				"   * Starfruit produces Artisan Goods that have ",
				"     some of the highe                          ",
			},
		},
		{
			name:       "Multiline and padding of note with 263 characters",
			input:      "Starfruit produces Artisan Goods that have some of the highest sell values in the game. Starfruit produces Artisan Goods that have some of the highest sell values in the game. Starfruit produces Artisan Goods that have some of the highest sell values in the game.",
			lineLength: 42,
			expected: []string{
				"   * Starfruit produces Artisan Goods that have ",
				"     some of the highest sell values in the     ",
				"     game. Starfruit produces Artisan Goods     ",
				"     that have some of the highest sell values  ",
				"     in the game. Starfruit produces Artisan    ",
				"     Goods that have some of the highest sell   ",
				"     values in the game.                        ",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := output.LineSplit(tc.input, tc.lineLength)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestLineBreaks(t *testing.T) {
	testCases := []struct {
		name       string
		input      string
		lineLength int
		expected   []string
	}{
		{
			name:       "Line length is over input length",
			input:      "xxxxxx xxxxxx",
			lineLength: 14,
			expected:   []string{"xxxxxx xxxxxx"},
		},
		{
			name:       "Line break at natural break following word",
			input:      "xxxxxx xxxxxx",
			lineLength: 6,
			expected:   []string{"xxxxxx", "xxxxxx"},
		},
		{
			name:       "Line break at natural break following space",
			input:      "xxxxxx xxxxxx",
			lineLength: 7,
			expected:   []string{"xxxxxx", "xxxxxx"},
		},
		{
			name:       "Break in the middle of a two word input",
			input:      "abcdef ghijkl",
			lineLength: 10,
			expected:   []string{"abcdef", "ghijkl"},
		},
		{
			name:       "Break in the middle of a short word",
			input:      "game. Starfruit produces Artisan Goods that have some of the highest sell values in",
			lineLength: 42,
			expected:   []string{"game. Starfruit produces Artisan Goods", "that have some of the highest sell values", "in"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := output.LineBreaks(tc.input, tc.lineLength)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
