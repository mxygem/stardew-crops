package output_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/jaysonesmith/stardew-crops/data"
	"github.com/jaysonesmith/stardew-crops/output"
	"github.com/stretchr/testify/assert"
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
					Continual:   false,
				},
				SeedPrices: data.Prices{GeneralStore: 40},
				Recipes:    []string{"Escargot", "Fiddlehead Risotto", "Oil of Garlic"},
				Notes:      []string{"Only available starting in year 2"},
			},
			f:        "raw",
			expected: bytes.NewBuffer([]byte(`{"name":"garlic","info":{"description":"Adds a wonderful zestiness to dishes. High quality garlic can be pretty spicy.","seed":"garlic seeds","growth_time":4,"season":["spring"],"continual":false,"trellis":false},"seed_prices":{"general_store":40,"jojamart":0},"recipes":["Escargot","Fiddlehead Risotto","Oil of Garlic"],"notes":["Only available starting in year 2"]}`)),
		},
		// {
		// 	name: "Pretty - Info response CropData Starfruit",
		// 	data: data.Crop{
		// 		Name: "starfruit",
		// 		Info: data.Info{
		// 			Description: "An extremely juicy fruit that grows in hot, humid weather. Slightly sweet with a sour undertone.",
		// 			Seed:        "starfruit seeds",
		// 			GrowthTime:  13,
		// 			Season:      []string{"summer"},
		// 			Continual:   false,
		// 		},
		// 		SeedPrices: data.Prices{Oasis: 400},
		// 		Quests:     []string{"A Soldier's Star"},
		// 		Notes:      []string{"Starfruit produces Artisan Goods that have some of the highest sell values in the game.", "Starfruit is required to build a Junimo Hut, purchased from the Wizard's Tower."},
		// 	},
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := output.Format(tc.data, tc.f)
			fmt.Println(actual.String())

			assert.Equal(t, tc.expected, actual)
		})
	}
}
