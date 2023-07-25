package routes

import (
	"Booking-Ticket-App/controller"
	"github.com/gin-gonic/gin"
)

type OrderRoute struct {
	orderController controller.OrderController
}

func NewOrderRoute(orderController controller.OrderController) OrderRoute {
	return OrderRoute{orderController}
}

func (r *OrderRoute) OrderRouteRegister(rt *gin.Engine) {
	router := rt.Group("/api/v1/orders")
	router.GET("/", r.orderController.GetAllOrderByDate)
	router.GET("/:id", r.orderController.GetOrderById)
	router.POST("/", r.orderController.CreateOrders)

}
