package routes

import (
	"Booking-Ticket-App/controller"
	"Booking-Ticket-App/services"
	"github.com/gin-gonic/gin"
)

type OrderRoute struct {
	orderController controller.OrderController
}

func NewOrderRoute(orderController controller.OrderController) OrderRoute {
	return OrderRoute{orderController}
}

func (r *OrderRoute) OrderRouteRegister(rt *gin.Engine, service services.UserService) {
	router := rt.Group("/api/v1/orders")
	router.POST("/", r.orderController.CreateOrders)

}
