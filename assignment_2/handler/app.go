package handler

import (
	"ddd/infra/database"
	"ddd/repository/order_repository/order_pg"
	"ddd/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	orderRepo := order_pg.NewOrderPG(db)

	orderService := service.NewOrderService(orderRepo)

	orderHandler := NewOrderHandler(orderService)

	r := gin.Default()

	r.POST("/orders", orderHandler.CreateOrder)
	r.PUT("/orders/:orderId", orderHandler.UpdateOrder)
	r.GET("/orders", orderHandler.GetOrders)
	r.DELETE("/orders/:orderId", orderHandler.DeleteOrder)

	r.Run(":8080")

}
