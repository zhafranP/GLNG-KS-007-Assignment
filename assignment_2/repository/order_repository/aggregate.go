package order_repository

import "ddd/entity"

type OrderItemMapped struct {
	Order entity.Order
	Items []entity.Item
}

type OrderItem struct {
	Order entity.Order
	Item  entity.Item
}

func (oim *OrderItemMapped) HandleMappingOrderWithItems(orderItem []OrderItem) []OrderItemMapped {
	ordersItemMapped := []OrderItemMapped{}

	for _, each := range orderItem {

		isOrderExist := false

		for i := range ordersItemMapped {
			if each.Order.OrderId == ordersItemMapped[i].Order.OrderId {
				isOrderExist = true
				ordersItemMapped[i].Items = append(ordersItemMapped[i].Items, each.Item)
				break
			}

		}

		// if each.Item.ItemID == 0 {
		// 	break
		// }

		if isOrderExist == false {
			orderItemMapped := OrderItemMapped{
				Order: each.Order,
			}

			orderItemMapped.Items = append(orderItemMapped.Items, each.Item)
			ordersItemMapped = append(ordersItemMapped, orderItemMapped)
		}
	}

	return ordersItemMapped
}
