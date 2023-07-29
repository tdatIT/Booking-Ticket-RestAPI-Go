package main

import (
	handler2 "Booking-Ticket-App/api/handler"
	"Booking-Ticket-App/config"
	"Booking-Ticket-App/domain/model"
	"Booking-Ticket-App/usecase/movie"
	"Booking-Ticket-App/usecase/movie/repository"
	"Booking-Ticket-App/usecase/movie_schedule"
	repository3 "Booking-Ticket-App/usecase/movie_schedule/repository"
	"Booking-Ticket-App/usecase/order"
	"Booking-Ticket-App/usecase/order/producer"
	"Booking-Ticket-App/usecase/user"
	repository2 "Booking-Ticket-App/usecase/user/repository"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

var (
	server           *gin.Engine
	ctx              context.Context
	redis_client     *redis.Client
	mongoClient      *mongo.Client
	kafka_connection *kafka.Conn
	//  Add the movie usecase and handle
	movieRepository model.MovieRepository
	movieUsecase    model.MovieUsecase
	movieCaching    model.CachingMovieRepository
	// Add the user usecase and handle
	userRepository model.UserRepository
	userUsecase    model.UserUsecase
	// Add the movie schedule api
	msRepository model.MovieScheduleRepository
	msUsecase    model.MovieScheduleUsecase
	// Add the order api
	messageProducer model.OrderDeliveryMessage
	orderUsecase    model.OrderUsecase
)

func main() {

	fmt.Println("Init web app")
	InitServer()
	defer config.CloseConnectionKafka(kafka_connection)
	err := server.Run(":5000")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Server is running on port 5000")
}
func InitServer() {
	timeOutContext := 10 * time.Second
	ctx = context.Background()

	ctx_timeout, _ := context.WithTimeout(ctx, 5*time.Second)
	mongoClient = config.ConnectDB(ctx_timeout)
	kafka_connection = config.ConnectionToKafka(ctx)
	redis_client = config.GetConnectToRedis()

	server = gin.Default()

	//Handle movie api movie endpoint
	movieRepository = repository.NewMongodbMovieRepository(mongoClient)
	movieCaching = repository.NewRedisCachingRepo(redis_client)
	movieUsecase = movie.NewMovieUseCase(movieRepository, movieCaching, timeOutContext)
	handler2.NewMovieController(movieUsecase, server)
	//Handle user api endpoint
	userRepository = repository2.NewMongoDBUserRepository(mongoClient)
	userUsecase = user.NewUserUsecase(userRepository, timeOutContext)
	handler2.NewUserHandler(userUsecase, server)
	//Handle movie schedule api
	msRepository = repository3.NewMongoDBMovieScheduleRepository(mongoClient)
	msUsecase = movie_schedule.NewMovieScheduleUsecase(msRepository, timeOutContext)
	handler2.NewMovieScheduleHandler(msUsecase, userRepository, server)
	//Handle order usecase api
	messageProducer = producer.NewKafkaProducer(kafka_connection)
	orderUsecase = order.NewOrderUsecase(messageProducer)
	handler2.NewOrderHandler(orderUsecase, server)
}
