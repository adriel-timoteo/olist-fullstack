package entity

import "time"

type DeliveredOrderTrend struct {
	DeliveryTime        time.Time
	TotalDeliveredCount int
}

type OrderStatusTrend struct {
	OrderTime  time.Time
	Status     string
	OrderCount int
}

type OrderByHour struct {
	Hour       int
	OrderCount float64
}
