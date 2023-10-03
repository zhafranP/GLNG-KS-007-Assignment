package service

import (
	"ddd/dto"
	"ddd/entity"
	"ddd/repository/order_repository"
	"net/http"
)

type orderService struct {
	OrderRepository order_repository.Repository
}

type OrderService interface {
	CreateOrder(newOrderRequest dto.NewOrderRequest) error
	GetOrders() (*dto.GetOrdersResponse, error)
	UpdateOrder(newOrderRequest dto.NewOrderRequest) (*dto.UpdateOrderResponse, error)
	DeleteOrder(orderId int) error
}

func NewOrderService(orderRepository order_repository.Repository) OrderService {
	return &orderService{
		OrderRepository: orderRepository,
	}
}

func (os *orderService) CreateOrder(newOrderRequest dto.NewOrderRequest) error {
	orderPayload := entity.Order{
		OrderedAt:    newOrderRequest.OrderedAt,
		CustomerName: newOrderRequest.CustomerName,
	}

	itemPayload := []entity.Item{}

	for _, each := range newOrderRequest.Items {
		item := entity.Item{
			ItemCode:    each.ItemCode,
			Description: each.Description,
			Quantity:    each.Quantity,
		}

		itemPayload = append(itemPayload, item)
	}

	err := os.OrderRepository.CreateOrder(orderPayload, itemPayload)

	if err != nil {
		return err
	}

	return nil
}

func (os *orderService) UpdateOrder(newOrderRequest dto.NewOrderRequest) (*dto.UpdateOrderResponse, error) {
	orderPayload := entity.Order{
		OrderId:      newOrderRequest.OrderID,
		OrderedAt:    newOrderRequest.OrderedAt,
		CustomerName: newOrderRequest.CustomerName,
	}

	itemPayload := []entity.Item{}

	for _, each := range newOrderRequest.Items {
		item := entity.Item{
			ItemCode:    each.ItemCode,
			Description: each.Description,
			Quantity:    each.Quantity,
		}

		itemPayload = append(itemPayload, item)
	}

	updatedOrder, err := os.OrderRepository.UpdateOrder(orderPayload, itemPayload)

	if err != nil {
		return nil, err
	}

	orderResult := dto.OrderWithItems{
		OrderId:      updatedOrder.Order.OrderId,
		CustomerName: updatedOrder.Order.CustomerName,
		OrderedAt:    updatedOrder.Order.OrderedAt,
		CreatedAt:    updatedOrder.Order.CreatedAt,
		UpdatedAt:    updatedOrder.Order.UpdatedAt,
		Items:        []dto.GetItemResponse{},
	}

	for _, eachItem := range updatedOrder.Items {
		item := dto.GetItemResponse{
			ItemId:      eachItem.ItemID,
			ItemCode:    eachItem.ItemCode,
			Quantity:    eachItem.Quantity,
			Description: eachItem.Description,
			OrderId:     eachItem.OrderId,
			CreatedAt:   eachItem.CreatedAt,
			UpdatedAt:   eachItem.UpdatedAt,
		}
		orderResult.Items = append(orderResult.Items, item)
	}

	response := dto.UpdateOrderResponse{
		StatusCode: http.StatusOK,
		Message:    "orders successfully fetched",
		Data:       orderResult,
	}

	return &response, nil

}

func (os *orderService) GetOrders() (*dto.GetOrdersResponse, error) {
	orders, err := os.OrderRepository.ReadOrders()

	if err != nil {
		return nil, err
	}

	orderResult := []dto.OrderWithItems{}

	for _, each := range orders {
		order := dto.OrderWithItems{
			OrderId:      each.Order.OrderId,
			CustomerName: each.Order.CustomerName,
			OrderedAt:    each.Order.OrderedAt,
			CreatedAt:    each.Order.CreatedAt,
			UpdatedAt:    each.Order.UpdatedAt,
			Items:        []dto.GetItemResponse{},
		}

		for _, eachItem := range each.Items {
			item := dto.GetItemResponse{
				ItemId:      eachItem.ItemID,
				ItemCode:    eachItem.ItemCode,
				Quantity:    eachItem.Quantity,
				Description: eachItem.Description,
				OrderId:     eachItem.OrderId,
				CreatedAt:   eachItem.CreatedAt,
				UpdatedAt:   eachItem.UpdatedAt,
			}
			order.Items = append(order.Items, item)
		}

		orderResult = append(orderResult, order)

	}

	response := dto.GetOrdersResponse{
		StatusCode: http.StatusOK,
		Message:    "orders successfully fetched",
		Data:       orderResult,
	}

	return &response, nil

}

func (os *orderService) DeleteOrder(orderId int) error {
	err := os.OrderRepository.DeleteOrder(orderId)

	if err != nil {
		return err
	}

	return nil
}
