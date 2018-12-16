package processors_test

import (
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
			expected: []string{"blueberry", "hot peppers"},
			err:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := processors.Search(tc.flags)

			for i, c := range actual {
				assert.Equal(t, tc.expected[i], c.Name)
			}
			assert.Equal(t, tc.err, err)
		})
	}
}
