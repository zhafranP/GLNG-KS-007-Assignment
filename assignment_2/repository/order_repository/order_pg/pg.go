package order_pg

import (
	"database/sql"
	"ddd/entity"
	"ddd/repository/order_repository"
	"fmt"
)

type orderPG struct {
	db *sql.DB
}

const (
	createOrderQuery = `
	INSERT INTO "orders"
	("ordered_at", "customer_name")
	VALUES ($1, $2)

	RETURNING "order_id"
	`

	createItemQuery = `
		INSERT INTO "items"
		("item_code", "description", "quantity", "order_id")
		VALUES ($1, $2, $3, $4)
	`

	getOrdersWithItemsQuery = `
	SELECT o.order_id, o.created_at, o.updated_at, o.customer_name, 
	i.item_id, i.created_at, i.updated_at, i.item_code, i.description, i.quantity,i.order_id
	FROM orders AS o
	LEFT JOIN items AS i ON o.order_id = i.order_id
	`

	updateOrders = `
	UPDATE orders 
	SET ordered_at = $1, customer_name = $2
	WHERE order_id = $3
	RETURNING order_id,created_at,updated_at,customer_name
	`

	updateItems = `
	UPDATE items 
	SET item_code = $1, description = $2,
	quantity = $3
	WHERE order_id = $4
	RETURNING item_id,created_at,updated_at,item_code,description,quantity,order_id
	`

	deleteOrders = `
	DELETE FROM orders
	WHERE order_id = $1
	`

	deleteItems = `
	DELETE FROM items
	WHERE order_id = $1
	`
)

func NewOrderPG(db *sql.DB) order_repository.Repository {
	return &orderPG{db: db}
}

func (orderPG *orderPG) CreateOrder(orderPayload entity.Order, itemPayload []entity.Item) error {
	tx, err := orderPG.db.Begin()

	if err != nil {
		return err
	}

	var orderId int
	orderRow := tx.QueryRow(createOrderQuery, orderPayload.OrderedAt, orderPayload.CustomerName)
	err = orderRow.Scan(&orderId)

	if err != nil {
		tx.Rollback()
		return err
	}

	for _, each := range itemPayload {

		_, err := tx.Exec(createItemQuery, each.ItemCode, each.Description, each.Quantity, orderId)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil

}

func (orderPG *orderPG) ReadOrders() ([]order_repository.OrderItemMapped, error) {
	rows, err := orderPG.db.Query(getOrdersWithItemsQuery)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	orderItems := []order_repository.OrderItem{}

	for rows.Next() {
		var orderItem order_repository.OrderItem

		err := rows.Scan(
			&orderItem.Order.OrderId, &orderItem.Order.CreatedAt, &orderItem.Order.UpdatedAt,
			&orderItem.Order.CustomerName, &orderItem.Item.ItemID, &orderItem.Item.CreatedAt,
			&orderItem.Item.UpdatedAt, &orderItem.Item.ItemCode, &orderItem.Item.Description,
			&orderItem.Item.Quantity, &orderItem.Item.OrderId,
		)

		if err != nil {
			return nil, err
		}

		orderItems = append(orderItems, orderItem)
	}

	fmt.Printf("%+v\n", orderItems)

	var result order_repository.OrderItemMapped
	return result.HandleMappingOrderWithItems(orderItems), nil
}

func (orderPG *orderPG) UpdateOrder(orderPayload entity.Order, itemsPayload []entity.Item) (*order_repository.OrderItemMapped, error) {
	tx, err := orderPG.db.Begin()

	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", itemsPayload)

	orderItem := order_repository.OrderItemMapped{}
	item := entity.Item{}

	err = tx.QueryRow(updateOrders, orderPayload.OrderedAt, orderPayload.CustomerName, orderPayload.OrderId).Scan(
		&orderItem.Order.OrderId, &orderItem.Order.CreatedAt, &orderItem.Order.UpdatedAt, &orderItem.Order.CustomerName,
	)

	fmt.Printf("%+v\n", orderItem)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, eachItem := range itemsPayload {
		// item_id,created_at,updated_at,item_code,description,quantity,order_id
		fmt.Printf("%+v\n", eachItem)
		err = tx.QueryRow(updateItems, eachItem.ItemCode, eachItem.Description, eachItem.Quantity, orderPayload.OrderId).Scan(
			&item.ItemID, &item.CreatedAt, &item.UpdatedAt,
			&item.ItemCode, &item.Description, &item.Quantity,
			&item.OrderId,
		)

		fmt.Println("err2", err)

		if err != nil {
			tx.Rollback()
			return nil, err
		}

		orderItem.Items = append(orderItem.Items, item)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &orderItem, nil

}

func (orderPG *orderPG) DeleteOrder(orderId int) error {
	tx, err := orderPG.db.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec(deleteOrders, orderId)

	if err != nil {
		tx.Rollback()
		return err
	}

	// _, err = tx.Exec(deleteItems, orderId)
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}
