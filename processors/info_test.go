package processors_test

import (
	"fmt"
	"testing"

	"github.com/jaysonesmith/stardew-crops/data"
	"github.com/jaysonesmith/stardew-crops/processors"
	"github.com/stretchr/testify/assert"
)

func TestInfo(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected data.Crop
		err      error
	}{
		{
			name:     "Empty args gets an error",
			input:    "",
			expected: data.Crop{},
			err:      fmt.Errorf("No crop name specified, please try again"),
		},
		{
			name:     "Unmatched crop returns an error",
			input:    "banana",
			expected: data.Crop{},
			err:      fmt.Errorf("Unable to find matching crop for banana"),
		},
		{
			name:  "Matched crop gets its data returned",
			input: "garlic",
			expected: data.Crop{
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
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			actual, err := processors.Info(tc.input)

			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.err, err)
		})
	}
}
