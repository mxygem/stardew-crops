package processors_test

import (
	"testing"

	"github.com/jaysonesmith/stardew-crops/processors"

	"github.com/stretchr/testify/assert"
)

func TestCropsByAllSeasons(t *testing.T) {
	expected := processors.CropsBySeasons{
		Spring: []string{"garlic", "potato"},
		Summer: []string{"blueberry", "radish", "starfruit"},
		Fall:   []string{"cranberries", "yams"},
	}

	actual := processors.CropsByAllSeasons()

	assert.Equal(t, expected, actual)
}
