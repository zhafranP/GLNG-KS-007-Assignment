package dto

import (
	"time"
)

// request

type NewOrderRequest struct {
	OrderID      int       `json:"order_id"`
	OrderedAt    time.Time `json:"ordered_at"`
	CustomerName string    `json:"customer_name"`
	Items        []NewItemRequest
}

// response

type GetItemResponse struct {
	ItemId      int       `json:"itemId"`
	ItemCode    string    `json:"itemCode"`
	Quantity    int       `json:"quantity"`
	Description string    `json:"description"`
	OrderId     int       `json:"orderId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type OrderWithItems struct {
	OrderId      int               `json:"orderId"`
	CustomerName string            `json:"customerName"`
	OrderedAt    time.Time         `json:"orderedAt"`
	CreatedAt    time.Time         `json:"createdAt"`
	UpdatedAt    time.Time         `json:"updatedAt"`
	Items        []GetItemResponse `json:"items"`
}

type UpdateOrderResponse struct {
	StatusCode int            `json:"statusCode"`
	Message    string         `json:"message"`
	Data       OrderWithItems `json:"data"`
}

type GetOrdersResponse struct {
	StatusCode int              `json:"statusCode"`
	Message    string           `json:"message"`
	Data       []OrderWithItems `json:"data"`
}

type DeleteOrderResponse struct {
	Message string `json:"message"`
}
