package types

// Weather API
type weatherContent struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type weatherMetrics struct {
	Temperature float32 `json:"temp"`
	Feelslike   float32 `json:"feels_like"`
	Humidity    float32 `json:"humidity"`
}
type wind struct {
	Speed float32 `json:"speed"`
}

type Weather struct {
	Weather []weatherContent `json:"weather"`
	Metrics weatherMetrics   `json:"main"`
	Wind    wind             `json:"wind"`
}
