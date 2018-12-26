package data

type CropData struct {
	Crops []Crop `json:"crops"`
}

// Crop represents the crop data model
type Crop struct {
	Name       string   `json:"name"`
	Info       Info     `json:"info"`
	SeedPrices Prices   `json:"seed_prices"`
	Bundles    []string `json:"bundles,omitempty"`
	Recipes    []string `json:"recipes,omitempty"`
	Quests     []string `json:"quests,omitempty"`
	Notes      []string `json:"notes,omitempty"`
}

type Info struct {
	Description string   `json:"description"`
	Seed        string   `json:"seed"`
	GrowthTime  int64    `json:"growth_time"`
	Season      []string `json:"season"`
	Continual   bool     `json:"continual,omitempty"`
	Regrowth    int64    `json:"regrowth,omitempty"`
	Trellis     bool     `json:"trellis,omitempty"`
}

type Prices struct {
	GeneralStore  int64               `json:"general_store,omitempty"`
	JojaMart      int64               `json:"jojamart,omitempty"`
	TravelingCart TravelingCartPrices `json:"traveling_cart,omitempty"`
	Oasis         int64               `json:"oasis,omitempty"`
	EggFestival   int64               `json:"egg_festival,omitempty"`
}

type TravelingCartPrices struct {
	Min int64 `json:"min,omitempty"`
	Max int64 `json:"max,omitempty"`
}
