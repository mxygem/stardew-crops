package processors_test

import (
	"testing"

	"github.com/jaysonesmith/stardew-crops/processors"
	"github.com/stretchr/testify/assert"
)

func TestSeason(t *testing.T) {
	testCases := []struct {
		name     string
		args     string
		expected processors.CropsBySeasons
	}{
		{
			name: "No args returns all crops",
			args: "",
			expected: processors.CropsBySeasons{
				Spring: []string{"garlic", "potato"},
				Summer: []string{"blueberry", "radish", "starfruit"},
				Fall:   []string{"cranberries", "yams"},
			},
		},
		{
			name: "Only spring crops",
			args: "spring",
			expected: processors.CropsBySeasons{
				Spring: []string{"garlic", "potato"},
			},
		},
		{
			name: "Only summer crops",
			args: "summer",
			expected: processors.CropsBySeasons{
				Summer: []string{"blueberry", "radish", "starfruit"},
			},
		},
		{
			name: "Only fall crops",
			args: "fall",
			expected: processors.CropsBySeasons{
				Fall: []string{"cranberries", "yams"},
			},
		},
		{
			name:     "Unmatched season",
			args:     "breakfast",
			expected: processors.CropsBySeasons{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := processors.Season(tc.args)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
