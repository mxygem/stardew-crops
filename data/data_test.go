package data_test

import (
	"testing"

	"github.com/jaysonesmith/stardew-crops/data"
	"github.com/stretchr/testify/assert"
)

func TestInitData(t *testing.T) {
	cropData := data.InitData()

	var cb, r, hp, s, g data.Crop
	for _, crop := range cropData.Crops {
		switch crop.Name {
		case "coffee bean":
			cb = crop
		case "rhubarb":
			r = crop
		case "hot pepper":
			hp = crop
		case "starfruit":
			s = crop
		case "grape":
			g = crop
		}
	}

	// spot check all fields across multiple crops
	// name
	assert.Equal(t, "coffee bean", cb.Name)
	// description
	assert.Equal(t, "The stalks are extremely tart, but make a great dessert when sweetened.", r.Info.Description)
	// seed
	assert.Equal(t, "grape starter", g.Info.Seed)
	// growth time
	assert.Equal(t, int64(13), s.Info.GrowthTime)
	// season
	assert.Equal(t, []string{"spring", "summer"}, cb.Info.Season)
	// continual
	assert.True(t, hp.Info.Continual)
	// regrowth
	assert.Equal(t, int64(3), hp.Info.Regrowth)
	// trellis
	assert.False(t, cb.Info.Trellis)
	assert.True(t, g.Info.Trellis)
	// prices general store
	assert.Equal(t, int64(60), g.SeedPrices.GeneralStore)
	// prices joja
	assert.Equal(t, int64(70), g.SeedPrices.JojaMart)
	// prices oasis
	assert.Equal(t, int64(400), s.SeedPrices.Oasis)
	// prices general nil
	assert.Equal(t, int64(0), s.SeedPrices.GeneralStore)
	// prices traveling cart
	assert.Equal(t, int64(100), cb.SeedPrices.TravelingCart.Min)
	assert.Equal(t, int64(2500), cb.SeedPrices.TravelingCart.Max)
	// bundles
	assert.Equal(t, []string{"Summer Foraging"}, g.Bundles)
	// recipes
	assert.Equal(t, []string{"Pepper Poppers", "Spicy Eel"}, hp.Recipes)
	// quests
	assert.Nil(t, r.Quests)
	assert.Equal(t, []string{"A Soldier's Star"}, s.Quests)
	// notes
	assert.Equal(t, []string{"Has a 1% chance to drop from dust sprite"}, cb.Notes)
}
