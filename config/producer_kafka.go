package config

import (
	"Booking-Ticket-App/domain/message"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func ConnectionToKafka(ctx context.Context) *kafka.Conn {

	conn, err := kafka.DialLeader(ctx, "tcp", "localhost:9092", TOPIC, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	return conn
}
func SendMessageToKafka(conn *kafka.Conn, msg message.OrderMessage) error {
	//encode object to byte
	json_message, err := json.Marshal(msg)
	if err != nil {
		log.Print(fmt.Errorf("could not encode order obj to message cause: %w", err))
		return err
	}
	//set timeout for write message
	err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	//send message
	_, err = conn.WriteMessages(
		kafka.Message{Value: json_message},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
		return err
	}
	return nil
}
func CloseConnectionKafka(conn *kafka.Conn) {
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}
