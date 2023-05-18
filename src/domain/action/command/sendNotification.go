package command

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
)

type SendNotification struct {
	notificationRepo model.NotificationRepository
}

func NewSendNotification(notificationRepo model.NotificationRepository) *SendNotification {
	return &SendNotification{
		notificationRepo: notificationRepo,
	}
}

func (c SendNotification) Do(ctx context.Context, notification model.Notification) error {
	return c.notificationRepo.Send(ctx, notification)
}
