package usecase

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/action/command"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
)

type CreateOrder struct {
	baseLogger            model.Logger
	findCustomerByIdQuery query.FindCustomerById
	createOrderCmd        command.CreateOrder
}

func NewCreateOrder(baseLogger model.Logger, findCustomerByIdQuery query.FindCustomerById, createOrderCmd command.CreateOrder) *CreateOrder {
	return &CreateOrder{
		baseLogger:            baseLogger.WithFields(model.LoggerFields{"useCase": "CreateOrder"}),
		findCustomerByIdQuery: findCustomerByIdQuery,
		createOrderCmd:        createOrderCmd,
	}
}

// Do : If it is successful, returns orderId
func (u *CreateOrder) Do(ctx context.Context, order model.Order) (int64, error) {
	log := u.baseLogger.WithRequestId(ctx).WithFields(model.LoggerFields{"order": order})

	if _, err := u.findCustomerByIdQuery.Do(ctx, order.CustomerId); err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Errorf("error when find customer by id %v", order.CustomerId)
		return 0, err
	}

	orderId, err := u.createOrderCmd.Do(ctx, order)
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("error when create order")
		return 0, err
	}

	log.Info("successfully created order")
	return orderId, nil
}
