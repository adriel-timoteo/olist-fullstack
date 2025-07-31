package dto

type CustomerCityCount struct {
	City  string `json:"city"`
	Count int    `json:"count"`
}

type Count struct {
	Count float64 `json:"count"`
}

type Rate struct {
	Rate float64 `json:"rate"`
}
