package processors_test

import (
	"testing"

	"github.com/jaysonesmith/stardew-crops/processors"

	"github.com/stretchr/testify/assert"
)

func TestCropsByAllSeasons(t *testing.T) {
	expected := map[string][]string{
		"spring": []string{"garlic", "potato"},
		"summer": []string{"blueberry", "radish", "starfruit"},
		"fall":   []string{"cranberries", "yams"},
	}

	actual := processors.CropsByAllSeasons()

	assert.Equal(t, expected, actual)
}
