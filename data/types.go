package data

type CropData struct {
	Crops []Crop `json:"crops"`
}

// Crop represents the crop data model
type Crop struct {
	Name       string   `json:"name"`
	Info       info     `json:"info"`
	SeedPrices prices   `json:"seed_prices"`
	Bundles    []string `json:"bundles,omitempty"`
	Recipes    []string `json:"recipes,omitempty"`
	Quests     []string `json:"quests,omitempty"`
	Notes      []string `json:"notes,omitempty"`
}

type info struct {
	Description string   `json:"description"`
	Seed        string   `json:"seed"`
	GrowthTime  int64    `json:"growth_time"`
	Season      []string `json:"season"`
	Continual   bool     `json:"continual"`
	Regrowth    int64    `json:"regrowth,omitempty"`
	Trellis     bool     `json:"trellis,omitempty"`
}

type prices struct {
	GeneralStore  int64               `json:"general_store,omitempty"`
	JojaMart      int64               `json:"jojamart,omitempty"`
	TravelingCart travelingCartPrices `json:"traveling_cart,omitempty"`
	Oasis         int64               `json:"Oasis,omitempty"`
}

type travelingCartPrices struct {
	Min int64 `json:"min"`
	Max int64 `json:"max"`
}
