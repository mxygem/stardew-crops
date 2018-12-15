package processors_test

import (
	"testing"

	"github.com/jaysonesmith/stardew-crops/processors"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		name     string
		flags    map[string]interface{}
		expected []string
		err      error
	}{
		{
			name:     "bundle search for Summer Crops returns two crops",
			flags:    map[string]interface{}{"bundle": "Summer Crops"},
			expected: []string{"blueberry", "hot pepper"},
			err:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := processors.Search(tc.flags)

			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.err, err)
		})
	}
}
