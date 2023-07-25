package main

import (
	"Booking-Ticket-App/config"
	"Booking-Ticket-App/controller"
	"Booking-Ticket-App/routes"
	"Booking-Ticket-App/services"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoClient *mongo.Client

	//  Add the Movies Service, Controllers and Routes
	movieService    services.MovieServices
	movieController controller.MovieController
	movieRoute      routes.MovieRoute
	//  Add the Movie Schedule Service, Controllers and Routes
	movieScheduleServices   services.MovieScheduleService
	movieScheduleController controller.MovieScheduleController
	movieScheduleRoute      routes.MovieScheduleRoute
	//  Add the Movies Service, Controllers and Routes
	orderService    services.OrderService
	orderController controller.OrderController
	orderRoute      routes.OrderRoute
	// Add authentication
	userServices   services.UserService
	authServices   services.AuthService
	authController controller.AuthController
	authRoute      routes.AuthRouter
)

func main() {
	fmt.Println("Init web app")
	InitServer()

	err := server.Run(":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Server is running on port 8080")
}
func InitServer() {
	ctx = context.TODO()
	mongoClient = config.DB
	//Handle movie api movie endpoint
	movieService = services.NewMovieClient(mongoClient, ctx)
	movieController = controller.NewMovieController(movieService)
	movieRoute = routes.NewMovieRouter(movieController)
	//Handle movie schedule endpoint
	movieScheduleServices = services.NewMovieScheduleService(mongoClient, ctx)
	movieScheduleController = controller.NewMovieScheduleController(movieScheduleServices)
	movieScheduleRoute = routes.NewMovieScheduleRoute(movieScheduleController)
	//Handle movie api movie endpoint
	orderService = services.NewOrderService(mongoClient, ctx)
	orderController = controller.NewOrderController(orderService)
	orderRoute = routes.NewOrderRoute(orderController)
	//Authentication api
	userServices = services.NewUserService(ctx, mongoClient)
	authServices = services.NewAuthService(mongoClient, ctx)
	authController = controller.NewAuthController(authServices, userServices)
	authRoute = routes.NewAuthRouter(authController)
	//create gin instance
	server = gin.Default()
	//register route
	movieRoute.MovieRouteRegister(server, userServices)
	movieScheduleRoute.MovieScheduleRouteRegister(server, userServices)
	orderRoute.OrderRouteRegister(server)
	authRoute.AuthRouteRegister(server)
}
