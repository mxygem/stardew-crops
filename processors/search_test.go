package processors_test

import (
	"fmt"
	"testing"

	"github.com/jaysonesmith/stardew-crops/processors"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		name     string
		flags    map[string]string
		expected []string
		err      error
	}{
		{
			name:     "bundle search for Quality Crops",
			flags:    map[string]string{"bundle": "Quality Crops"},
			expected: []string{"corn"},
			err:      nil,
		},
		{
			name:     "bundle search for Summer Crops",
			flags:    map[string]string{"bundle": "Summer Crops"},
			expected: []string{"blueberry", "hot pepper"},
			err:      nil,
		},
		{
			name:     "bundle search for an unmatched bundle",
			flags:    map[string]string{"bundle": "Construction"},
			expected: []string{},
			err:      fmt.Errorf("No matching crops found"),
		},
		{
			name:     "bundle search with no value",
			flags:    map[string]string{"bundle": ""},
			expected: []string{},
			err:      fmt.Errorf("A value must be provided for the bundle flag"),
		},
		{
			name:     "growth search for 6 days",
			flags:    map[string]string{"growth": "6"},
			expected: []string{"potato", "radish"},
			err:      nil,
		},
		{
			name:     "growth search for 100 days",
			flags:    map[string]string{"growth": "100"},
			expected: []string{},
			err:      fmt.Errorf("No matching crops found"),
		},
		{
			name:     "growthgt search for 5 days",
			flags:    map[string]string{"growthgt": "5"},
			expected: []string{"coffee bean", "potato", "rhubarb", "blueberry", "corn", "radish", "hot pepper", "starfruit", "cranberries", "grape", "yam"},
			err:      nil,
		},
		{
			name:     "growthgt search for 13 days",
			flags:    map[string]string{"growthgt": "13"},
			expected: []string{"rhubarb", "blueberry", "corn", "starfruit"},
			err:      nil,
		},
		{
			name:     "growthgt search with no value",
			flags:    map[string]string{"growthgt": ""},
			expected: []string{},
			err:      fmt.Errorf("A value must be provided for the growthgt flag"),
		},
		{
			name:     "growthlt search for 5 days",
			flags:    map[string]string{"growthlt": "5"},
			expected: []string{"garlic", "hot pepper"},
			err:      nil,
		},
		{
			name:     "growthlt search for 13 days",
			flags:    map[string]string{"growthlt": "13"},
			expected: []string{"coffee bean", "garlic", "potato", "rhubarb", "blueberry", "radish", "hot pepper", "starfruit", "cranberries", "grape", "yam"},
			err:      nil,
		},
		{
			name:     "growthlt search with no value",
			flags:    map[string]string{"growthlt": ""},
			expected: []string{},
			err:      fmt.Errorf("A value must be provided for the growthlt flag"),
		},
		{
			name:     "season search for Summer",
			flags:    map[string]string{"season": "summer"},
			expected: []string{"coffee bean", "blueberry", "corn", "radish", "hot pepper", "starfruit"},
			err:      nil,
		},
		{
			name:     "season search for an unmatched season",
			flags:    map[string]string{"season": "breakfast"},
			expected: []string{},
			err:      fmt.Errorf("No matching crops found"),
		},
		{
			name:     "season search with no value",
			flags:    map[string]string{"season": ""},
			expected: []string{},
			err:      fmt.Errorf("A value must be provided for the season flag"),
		},
		{
			name:     "trellis search for crops with trellis",
			flags:    map[string]string{"trellis": "true"},
			expected: []string{"grape"},
			err:      nil,
		},
		{
			name:     "trellis search for crops without trellis",
			flags:    map[string]string{"trellis": "false"},
			expected: []string{"coffee bean", "garlic", "potato", "rhubarb", "blueberry", "corn", "radish", "hot pepper", "starfruit", "cranberries", "yam"},
			err:      nil,
		},
		{
			name:     "trellis search with no value",
			flags:    map[string]string{"trellis": ""},
			expected: []string{},
			err:      fmt.Errorf("A value must be provided for the trellis flag"),
		},
		{
			name:     "trellis search with non-bool value",
			flags:    map[string]string{"trellis": "abc"},
			expected: []string{},
			err:      fmt.Errorf("Please pass in true or false for the trellis flag"),
		},
		{
			name:     "continual search for multi harvest crops",
			flags:    map[string]string{"continuous": "true"},
			expected: []string{"coffee bean", "blueberry", "corn", "hot pepper", "cranberries", "grape"},
			err:      nil,
		},
		{
			name:     "continual search for single harvest crops",
			flags:    map[string]string{"continuous": "false"},
			expected: []string{"garlic", "potato", "rhubarb", "radish", "starfruit", "yam"},
			err:      nil,
		},
		{
			name:     "continuous search with no value",
			flags:    map[string]string{"continuous": ""},
			expected: []string{},
			err:      fmt.Errorf("A value must be provided for the continuous flag"),
		},
		{
			name:     "continuous search with non-bool value",
			flags:    map[string]string{"continuous": "abc"},
			expected: []string{},
			err:      fmt.Errorf("Please pass in true or false for the continuous flag"),
		},
		{
			name:     "multi flag search with three values",
			flags:    map[string]string{"growth": "13", "season": "summer", "continuous": "false"},
			expected: []string{"starfruit"},
			err:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := processors.Search(tc.flags)

			assert.Equal(t, tc.err, err)
			for i, c := range actual {
				assert.Equal(t, tc.expected[i], c.Name)
			}
		})
	}
}
