package processors_test

import (
	"fmt"
	"testing"

	"github.com/jaysonesmith/stardew-crops/processors"
	"github.com/stretchr/testify/assert"
)

func TestSeason(t *testing.T) {
	testCases := []struct {
		name     string
		args     string
		expected processors.CropsBySeasons
		err      error
	}{
		{
			name: "No args returns all crops",
			args: "",
			expected: processors.CropsBySeasons{
				Spring: []string{"coffee bean", "garlic", "potato", "rhubarb"},
				Summer: []string{"coffee bean", "blueberry", "corn", "radish", "starfruit"},
				Fall:   []string{"corn", "cranberries", "yam"},
			},
			err: nil,
		},
		{
			name: "Only spring crops",
			args: "spring",
			expected: processors.CropsBySeasons{
				Spring: []string{"coffee bean", "garlic", "potato", "rhubarb"},
			},
			err: nil,
		},
		{
			name: "Only summer crops",
			args: "summer",
			expected: processors.CropsBySeasons{
				Summer: []string{"coffee bean", "blueberry", "corn", "radish", "starfruit"},
			},
			err: nil,
		},
		{
			name: "Only fall crops",
			args: "fall",
			expected: processors.CropsBySeasons{
				Fall: []string{"corn", "cranberries", "yam"},
			},
			err: nil,
		},
		{
			name:     "Unmatched season",
			args:     "breakfast",
			expected: processors.CropsBySeasons{},
			err:      fmt.Errorf("Unknown season for breakfast"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := processors.Season(tc.args)

			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.err, err)
		})
	}
}
