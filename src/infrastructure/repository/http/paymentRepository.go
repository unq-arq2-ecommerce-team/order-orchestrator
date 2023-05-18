package http

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/infrastructure/config"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/infrastructure/logger"
	"net/http"
)

type paymentRepository struct {
	logger         model.Logger
	client         *http.Client
	makePaymentUrl string
}

func NewPaymentRepository(baseLogger model.Logger, client *http.Client, paymentEndpoints *config.PaymentEndpoints) model.PaymentRepository {
	return paymentRepository{
		logger:         baseLogger.WithFields(logger.Fields{"logger": "http.paymentRepository"}),
		client:         client,
		makePaymentUrl: paymentEndpoints.MakePayment,
	}
}

func (repo paymentRepository) MakePayment(ctx context.Context, payment *model.Payment) error {
	url := repo.makePaymentUrl
	log := repo.logger.WithFields(logger.Fields{"makePaymentUrl": url, "payment": payment})
	log.Debugf("making order payment...")

	log = repo.logger.WithFields(logger.Fields{"url": url})

	res, err := MakeAndDoRequest(ctx, log, repo.client, http.MethodPost, url, payment)
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
