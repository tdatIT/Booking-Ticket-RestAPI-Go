package message

import "time"

type OrderMessage struct {
	UserID          string    `json:"user_id"`
	MovieScheduleID string    `json:"movie_schedule_id"`
	Total           int       `json:"total"`
	Seats           []string  `json:"seats"`
	CreateAt        time.Time `json:"create_at"`
	SendBy          string    `json:"send_by"`
	TicketPrice     int       `json:"ticket_price"`
}
