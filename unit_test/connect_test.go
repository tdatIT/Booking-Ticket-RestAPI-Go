package unit_test

import (
	"Booking-Ticket-App/config"
	"fmt"
	"testing"
)

func TestConnection(t *testing.T) {
	DB := config.DB
	fmt.Println(config.GetCollection(DB, "users").Name())
}
