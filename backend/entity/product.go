package entity

import "time"

type DeliveredProductTrend struct {
	DeliveryTime        time.Time
	TotalDeliveredCount int
}

type ProductStatusTrend struct {
	OrderTime  time.Time
	Status     string
	OrderCount int
}

type ProductCategoryCount struct {
	Category     string
	ProductCount int
}
