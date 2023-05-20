package model

import (
	"context"
	"fmt"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/util"
)

const (
	CustomerRecipientType = "customer"
	SellerRecipientType   = "seller"
	eventPaymentOk        = "purchase_successful"
	channelEmail          = "email"
)

type (
	Notification struct {
		Channel   string    `json:"channel"`
		Event     Event     `json:"event"`
		Recipient Recipient `json:"recipient"`
	}

	Event struct {
		Name   string `json:"name"`
		Detail string `json:"detail"`
	}

	Recipient struct {
		Type string `json:"type"`
		Id   int64  `json:"id"`
	}
)

func NewEmailNotificationOrderPayed(recipientType string, userId int64, order Order) Notification {
	return Notification{
		Channel: channelEmail,
		Event: Event{
			Name:   eventPaymentOk,
			Detail: getDetailByUserType(recipientType, order),
		},
		Recipient: Recipient{
			Type: recipientType,
			Id:   userId,
		},
	}
}

func getDetailByUserType(userType string, order Order) string {
	switch userType {
	case CustomerRecipientType:
		return fmt.Sprintf("%s - $%v - Numero de orden: #%v", order.Product.Name, order.Product.Price, order.Id)
	default:
		return fmt.Sprintf("%s - $%v", order.Product.Name, order.Product.Price)
	}
}

func (n *Notification) String() string {
	return util.ParseStruct("Notification", n)
}

//go:generate mockgen -destination=../mock/notificationRepository.go -package=mock -source=notification.go
type NotificationRepository interface {
	Send(ctx context.Context, notification Notification) error
}
