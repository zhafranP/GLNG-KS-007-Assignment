package entity

import "time"

type Item struct {
	ItemID      int
	ItemCode    string
	Description string
	Quantity    int
	OrderId     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
