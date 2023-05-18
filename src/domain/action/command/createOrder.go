package command

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
)

type CreateOrder struct {
	orderRepo model.OrderRepository
}

func NewCreateOrder(orderRepo model.OrderRepository) *CreateOrder {
	return &CreateOrder{
		orderRepo: orderRepo,
	}
}

func (c CreateOrder) Do(ctx context.Context, order model.Order) (int64, error) {
	return c.orderRepo.Create(ctx, order)
}
