package handler

import (
	"ddd/dto"
	"ddd/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	OrderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) orderHandler {
	return orderHandler{
		OrderService: orderService,
	}
}

func (oh *orderHandler) CreateOrder(c *gin.Context) {
	var newOrderRequest dto.NewOrderRequest

	if err := c.ShouldBindJSON(&newOrderRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "invalid json request",
		})
		return
	}

	err := oh.OrderService.CreateOrder(newOrderRequest)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": "successfully creating a product",
	})

}

func (oh *orderHandler) GetOrders(ctx *gin.Context) {
	response, err := oh.OrderService.GetOrders()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (oh *orderHandler) UpdateOrder(c *gin.Context) {
	paramId := c.Param("orderId")
	var newOrderRequest dto.NewOrderRequest

	if err := c.ShouldBindJSON(&newOrderRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "invalid json request",
		})
		return
	}

	parse, err := strconv.ParseInt(paramId, 10, 0)
	newOrderRequest.OrderID = int(parse)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	response, err := oh.OrderService.UpdateOrder(newOrderRequest)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(response.StatusCode, response)
}

func (oh *orderHandler) DeleteOrder(c *gin.Context) {
	paramId := c.Param("orderId")
	parse, err := strconv.ParseInt(paramId, 10, 0)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	err = oh.OrderService.DeleteOrder(int(parse))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}
	c.JSON(http.StatusOK, dto.DeleteOrderResponse{Message: "Data Order Berhasil Dihapus"})
}
