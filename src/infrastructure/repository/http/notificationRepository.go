package http

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/infrastructure/config"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/infrastructure/logger"
	"net/http"
)

type notificationRepository struct {
	logger              model.Logger
	client              *http.Client
	sendNotificationUrl string
}

func NewNotificationRepository(baseLogger model.Logger, client *http.Client, notificationEndpoints *config.NotificationEndpoints) model.NotificationRepository {
	return notificationRepository{
		logger:              baseLogger.WithFields(logger.Fields{"logger": "http.NotificationRepository"}),
		client:              client,
		sendNotificationUrl: notificationEndpoints.Send,
	}
}

func (repo notificationRepository) Send(ctx context.Context, notification model.Notification) error {
	url := repo.sendNotificationUrl
	log := repo.logger.WithRequestId(ctx).WithFields(logger.Fields{"sendNotificationUrl": url, "notification": notification})
	log.Debugf("sending notification repository")

	log = repo.logger.WithFields(logger.Fields{"url": url})

	res, err := MakeAndDoRequest(ctx, log, repo.client, http.MethodPost, url, notification)
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("error when make and do request http")
		return err
	}
	statusCode := res.StatusCode
	if !IsStatusCode2XX(statusCode) {
		log.WithFields(model.LoggerFields{"error": err}).Error("error in status code")
		return NewUnexpectedError("order repository", statusCode, url)
	}
	return nil
}
