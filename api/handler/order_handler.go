package handler

import (
	model2 "Booking-Ticket-App/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler struct {
	orderUsecase model2.OrderUsecase
}

func NewOrderHandler(orderUsecase model2.OrderUsecase, ctx *gin.Engine) {
	handler := OrderHandler{orderUsecase: orderUsecase}
	router := ctx.Group("/api/v1/orders")
	router.POST("/", handler.CreateOrders)
}

func (ord *OrderHandler) CreateOrders(ctx *gin.Context) {
	var orderDTO *model2.OrderDTO
	if err := ctx.ShouldBindJSON(&orderDTO); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	err := ord.orderUsecase.CreateNewOrder(orderDTO)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "The order has been created successfully. The system is processing it."})
}
