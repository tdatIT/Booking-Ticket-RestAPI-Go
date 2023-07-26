package controller

import (
	model2 "Booking-Ticket-App/data/model"
	"Booking-Ticket-App/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) OrderController {
	return OrderController{orderService: orderService}
}

/*
	func (ord *OrderController) GetAllOrderByDate(ctx *gin.Context) {
		var orders []model2.Order
		var err error
		dateFilter := ctx.Query("date")
		orders, err = ord.orderService.GetAllOrderByDate(dateFilter)
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": orders})

}

	func (ord *OrderController) GetOrderById(ctx *gin.Context) {
		var order *model2.Order
		var err error
		orderId := ctx.Param("id")
		order, err = ord.orderService.GetOrderById(orderId)
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": order})

}
*/
func (ord *OrderController) CreateOrders(ctx *gin.Context) {
	var orderDTO *model2.OrderDTO
	if err := ctx.ShouldBindJSON(&orderDTO); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	err := ord.orderService.CreateNewOrder(orderDTO)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "The order has been created successfully. The system is processing it."})
}
