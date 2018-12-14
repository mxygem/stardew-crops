package processors_test

import (
	"testing"

	"github.com/jaysonesmith/stardew-crops/data"
	"github.com/jaysonesmith/stardew-crops/processors"
	"github.com/stretchr/testify/assert"
)

func TestAllCrops(t *testing.T) {
	actual := processors.AllCrops()

	assert.Equal(t, data.InitData(), actual)
}
