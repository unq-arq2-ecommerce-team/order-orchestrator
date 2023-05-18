package model

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/util"
)

const (
	CustomerRecipientType = "customer"
	SellerRecipientType   = "seller"
	eventPaymentOk        = "purchase_successful"
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

func NewNotificationOrderPayed(recipientType string, userId int64, detail string) Notification {
	return Notification{
		Channel: "email",
		Event: Event{
			Name:   eventPaymentOk,
			Detail: detail,
		},
		Recipient: Recipient{
			Type: recipientType,
			Id:   userId,
		},
	}
}

func (n *Notification) String() string {
	return util.ParseStruct("Notification", n)
}

//go:generate mockgen -destination=../mock/notificationRepository.go -package=mock -source=notification.go
type NotificationRepository interface {
	Send(ctx context.Context, notification Notification) error
}
