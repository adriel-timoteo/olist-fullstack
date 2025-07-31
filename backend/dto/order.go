package dto

type DeliveredTrend struct {
	Time  string `json:"time"`
	Count int    `json:"count"`
}

type OrderStatusSnapshot struct {
	Time   string `json:"time"`
	Status string `json:"status"`
	Count  int    `json:"count"`
}

type OrderByHour struct {
	Hour       int     `json:"hour"`
	OrderCount float64 `json:"count"`
}
