package order

import (
	"Booking-Ticket-App/domain/model"
)

type OrderUsecase struct {
	messageProducer model.OrderDeliveryMessage
}

func NewOrderUsecase(messageBroker model.OrderDeliveryMessage) OrderUsecase {
	return OrderUsecase{
		messageProducer: messageBroker,
	}
}

func (ord OrderUsecase) CreateNewOrder(dto *model.OrderDTO) error {
	return ord.messageProducer.SendMessageToConsumer(dto)
}
