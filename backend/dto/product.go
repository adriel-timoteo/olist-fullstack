package dto

type DeliveredTrend struct {
	DeliveryTime        string `json:"delivery_time"`
	TotalDeliveredCount int    `json:"total_delivered_products"`
}
