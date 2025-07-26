package dto

type DeliveredTrend struct {
	Time  string `json:"time"`
	Count int    `json:"count"`
}

type ProductStatusSnapshot struct {
	Time   string `json:"time"`
	Status string `json:"status"`
	Count  int    `json:"count"`
}

type ProductCategoryCount struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
}
