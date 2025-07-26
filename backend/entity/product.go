package entity

import "time"

type DeliveredProductTrend struct {
	DeliveryTime        time.Time
	TotalDeliveredCount int
}
