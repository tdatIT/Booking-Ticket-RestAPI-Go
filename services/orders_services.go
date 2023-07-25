package services

import (
	"Booking-Ticket-App/config"
	"Booking-Ticket-App/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type OrderService struct {
	client           *mongo.Client
	ticketCollection *mongo.Collection
	orderCollection  *mongo.Collection
	ctx              context.Context
}

func NewOrderService(client *mongo.Client, ctx context.Context) OrderService {
	return OrderService{
		client:           client,
		ticketCollection: config.GetCollection(client, "tickets"),
		orderCollection:  config.GetCollection(client, "orders"),
		ctx:              ctx,
	}
}
func (ord *OrderService) GetOrderById(orderId string) (*model.Order, error) {
	orderObjId, _ := primitive.ObjectIDFromHex(orderId)
	query := bson.M{"_id": orderObjId}
	var order *model.Order
	err := ord.orderCollection.FindOne(ord.ctx, query).Decode(&order)
	if err != nil {
		log.Print(fmt.Errorf("cannot get order cause: %w", err))
		return nil, err
	}
	return order, nil
}
func (ord *OrderService) GetAllOrderByDate(datetime string) ([]model.Order, error) {
	filter := bson.M{
		"$expr": bson.M{
			"$eq": []interface{}{
				bson.M{
					"$dateToString": bson.M{
						"format": "%Y-%m-%d",
						"date":   "$created_at",
					},
				},
				datetime,
			},
		},
	}
	orders := make([]model.Order, 0)
	rs, err := ord.orderCollection.Find(ord.ctx, filter)
	if err != nil {
		log.Print(fmt.Errorf("cannot get order by date cause: %w", err))
		return nil, err
	}
	if err = rs.All(ord.ctx, &orders); err != nil {
		log.Print(fmt.Errorf("cannot get order by date cause: %w", err))
		return nil, err
	}
	return orders, nil
}

func (ord *OrderService) CreateNewOrder(dto *model.OrderDTO) (string, error) {

	var result *mongo.InsertOneResult
	var newOrderId string
	var err error
	//insert tickets
	scheduleId, _ := primitive.ObjectIDFromHex(dto.MovieScheduleID)
	ticketIdArr := make([]primitive.ObjectID, 0)
	for _, ticketSeat := range dto.TicketSeats {
		ticket := model.Ticket{
			ID:         primitive.NewObjectID(),
			ShowtimeID: scheduleId,
			SeatNumber: ticketSeat,
			Price:      dto.PriceTicket,
			Status:     "sold",
			CreatedAt:  time.Now(),
		}
		result, err = ord.ticketCollection.InsertOne(ord.ctx, ticket)
		if err != nil {
			log.Print(fmt.Errorf("insert ticket was failed cause:%w", err))
			return "", err
		} else {
			ticketIdArr = append(ticketIdArr, result.InsertedID.(primitive.ObjectID))
		}
	}
	userId, _ := primitive.ObjectIDFromHex(dto.UserID)
	orderObj := model.Order{
		ID:        primitive.NewObjectID(),
		UserID:    userId,
		Total:     dto.Total,
		Status:    "success",
		CreatedAt: time.Now(),
		Tickets:   ticketIdArr,
	}
	if result, err := ord.orderCollection.InsertOne(ord.ctx, orderObj); err != nil {
		log.Print(fmt.Errorf("insert order was failed cause:%w", err))
	} else {
		newOrderId = result.InsertedID.(primitive.ObjectID).Hex()
	}
	return newOrderId, err
}
