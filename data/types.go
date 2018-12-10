package data

type CropData struct {
	Crops []Crop `json:"crops"`
}

// Crop represents the crop data model
type Crop struct {
	Name       string   `json:"name"`
	Info       info     `json:"info"`
	SeedPrices prices   `json:"seed_prices"`
	Bundles    []string `json:"bundles"`
	Recipes    []string `json:"recipes"`
	Quests     []string `json:"quests"`
	Notes      []string `json:"notes,omitempty"`
}

type info struct {
	Description string `json:"description"`
	Seed        string `json:"seed"`
	GrowthTime  int64  `json:"growth_time"`
	Season      string `json:"season"`
	Continual   bool   `json:"continual"`
}

type prices struct {
	GeneralStore  int64               `json:"general_store"`
	JojaMart      int64               `json:"jojamart"`
	TravelingCart travelingCartPrices `json:"traveling_cart"`
	NightMarket   int64               `json:"night_market"`
}

type travelingCartPrices struct {
	Min int64 `json:"min"`
	Max int64 `json:"max"`
}
