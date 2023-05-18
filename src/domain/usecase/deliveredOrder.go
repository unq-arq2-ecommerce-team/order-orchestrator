package usecase

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/action/command"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
)

type DeliveredOrder struct {
	baseLogger        model.Logger
	deliveredOrderCmd command.DeliveredOrder
}

func NewDeliveredOrder(baseLogger model.Logger, deliveredOrderCmd command.DeliveredOrder) *DeliveredOrder {
	return &DeliveredOrder{
		baseLogger:        baseLogger.WithFields(model.LoggerFields{"useCase": "DeliveredOrder"}),
		deliveredOrderCmd: deliveredOrderCmd,
	}
}

func (u DeliveredOrder) Do(ctx context.Context, orderId int64) error {
	log := u.baseLogger.WithFields(model.LoggerFields{"orderId": orderId})
	if err := u.deliveredOrderCmd.Do(ctx, orderId); err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("error when delivered order")
		return err
	}
	log.Info("successful order delivered")
	return nil
}
