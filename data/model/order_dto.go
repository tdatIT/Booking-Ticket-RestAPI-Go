package model

type OrderDTO struct {
	ID              string   `json:"id"`
	UserID          string   `json:"user_id"`
	Total           int64    `json:"total"`
	MovieScheduleID string   `json:"movie_schedule_id"`
	TicketSeats     []string `json:"tickets"`
	TicketPrice     int      `json:"ticket_price"`
}
