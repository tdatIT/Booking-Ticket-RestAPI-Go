package producer

import (
	"Booking-Ticket-App/config"
	"Booking-Ticket-App/domain/message"
	"Booking-Ticket-App/domain/model"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type KafkaProducer struct {
	kafka_cnt *kafka.Conn
}

func NewKafkaProducer(kafka *kafka.Conn) KafkaProducer {
	return KafkaProducer{
		kafka_cnt: kafka,
	}
}

func (mb KafkaProducer) SendMessageToConsumer(dto *model.OrderDTO) error {
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
	err := config.SendMessageToKafka(mb.kafka_cnt, orderMsg)
	if err != nil {
		return err
	}
	log.Printf("Send order message to kafka at %v", orderMsg.SendBy)
	return nil

}
