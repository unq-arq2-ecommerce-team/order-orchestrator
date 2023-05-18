package model

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/util"
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

func (n *Notification) String() string {
	return util.ParseStruct("Notification", n)
}

//go:generate mockgen -destination=../mock/notificationRepository.go -package=mock -source=notification.go
type NotificationRepository interface {
	Send(ctx context.Context, notification Notification) error
}
