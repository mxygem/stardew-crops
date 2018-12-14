package data_test

import (
	"testing"

	"github.com/jaysonesmith/stardew-crops/data"
	"github.com/stretchr/testify/assert"
)

func TestInitData(t *testing.T) {
	cropData := data.InitData()

	var cb, hp, s data.Crop
	for _, crop := range cropData.Crops {
		switch crop.Name {
		case "coffee bean":
			cb = crop
		case "hot pepper":
			hp = crop
		case "starfruit":
			s = crop
		}
	}

	assert.Equal(t, []string{"spring", "summer"}, cb.Info.Season)
	assert.False(t, cb.Info.Trellis)
	assert.Equal(t, int64(2500), cb.SeedPrices.TravelingCart.Max)

	assert.Equal(t, int64(3), hp.Info.Regrowth)
	assert.Equal(t, []string{"Pepper Poppers", "Spicy Eel"}, hp.Recipes)

	assert.Equal(t, int64(400), s.SeedPrices.Oasis)
}
