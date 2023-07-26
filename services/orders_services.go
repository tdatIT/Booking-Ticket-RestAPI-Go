package services

import (
	"Booking-Ticket-App/config"
	"Booking-Ticket-App/data/message"
	model2 "Booking-Ticket-App/data/model"
	"context"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type OrderService struct {
	client           *mongo.Client
	ticketCollection *mongo.Collection
	orderCollection  *mongo.Collection
	ctx              context.Context
	kafka_cnt        *kafka.Conn
}

func NewOrderService(client *mongo.Client, ctx context.Context, kafka *kafka.Conn) OrderService {
	return OrderService{
		client:           client,
		ticketCollection: config.GetCollection(client, "tickets"),
		orderCollection:  config.GetCollection(client, "orders"),
		ctx:              ctx,
		kafka_cnt:        kafka,
	}
}
func (ord *OrderService) CreateNewOrder(dto *model2.OrderDTO) error {
	//insert tickets
	orderMsg := message.OrderMessage{
		UserID:          dto.UserID,
		MovieScheduleID: dto.MovieScheduleID,
		Total:           int(dto.Total),
		Seats:           dto.TicketSeats,
		CreateAt:        time.Now(),
		TicketPrice:     dto.TicketPrice,
		SendBy:          "from:/api/v1/orders [VN-Go-BE-Service]",
	}
	err := config.SendMessageToKafka(ord.kafka_cnt, orderMsg)
	if err != nil {
		return err
	}
	return nil
}

/*func (ord *OrderService) GetOrderById(orderId string) (*model2.Order, error) {
	orderObjId, _ := primitive.ObjectIDFromHex(orderId)
	query := bson.M{"_id": orderObjId}
	var order *model2.Order
	err := ord.orderCollection.FindOne(ord.ctx, query).Decode(&order)
	if err != nil {
		log.Print(fmt.Errorf("cannot get order cause: %w", err))
		return nil, err
	}
	return order, nil
}
func (ord *OrderService) GetAllOrderByDate(datetime string) ([]model2.Order, error) {
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
	orders := make([]model2.Order, 0)
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
}*/
