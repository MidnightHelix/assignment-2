package model

type Item struct {
	ID          uint64 `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint64 `json:"quantity"`
	OrderID     uint64 `json:"order_id"`
	// CreatedAt    time.Time `json:"ordered_at" gorm:"column:ordered_at"`
}
