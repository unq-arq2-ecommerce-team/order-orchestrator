package usecase

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/action/command"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model/exception"
)

type PayOrder struct {
	baseLogger         model.Logger
	findOrderByIdQuery query.FindOrderById
	confirmOrderCmd    command.ConfirmOrder
}

func NewPayOrder(baseLogger model.Logger, findOrderByIdQuery query.FindOrderById, confirmOrderCmd command.ConfirmOrder) *PayOrder {
	return &PayOrder{
		baseLogger:         baseLogger.WithFields(model.LoggerFields{"useCase": "PayOrder"}),
		findOrderByIdQuery: findOrderByIdQuery,
		confirmOrderCmd:    confirmOrderCmd,
	}
}

// Do
//  1. [DONE] Find order and validate if could be payed
//  2. TODO: Pay order
//  3. [DONE] Confirm order
//  4. TODO: Notify seller and customer was purchase successfully
func (u *PayOrder) Do(ctx context.Context, orderId int64) error {
	log := u.baseLogger.WithFields(model.LoggerFields{"orderId": orderId})
	log.Debug("paying order ...")
	order, err := u.findOrderByIdQuery.Do(ctx, orderId)
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Errorf("error when find order by id %v", orderId)
		return err
	}
	log = log.WithFields(model.LoggerFields{"order": order})

	if order.WasPaid() {
		log.Errorf("cannot pay order because not apply, orderState: %s", order.State)
		return exception.OrderWasPaid{Id: orderId}
	}

	//TODO: Pay order (need error handling

	// If pay was successfull
	if err := u.confirmOrderCmd.Do(ctx, order.Id); err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Errorf("error when confirm order by id %v", orderId)
		return err
	}

	//TODO: Notify order (don't need error handling)

	return nil
}
