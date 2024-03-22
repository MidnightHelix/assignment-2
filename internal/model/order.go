package model

import (
	"time"
)

type Order struct {
	ID           uint64    `json:"order_id" example:"1"`
	CustomerName string    `json:"customer_name" example:"testing"`
	OrderedAt    time.Time `json:"ordered_at" example:"2019-11-10T04:21:46+07:00"`
	Items        []Item    `json:"items"`
}
