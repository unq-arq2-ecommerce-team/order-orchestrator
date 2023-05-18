package command

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
)

type ConfirmOrder struct {
	orderRepo model.OrderRepository
}

func NewConfirmOrder(orderRepo model.OrderRepository) *ConfirmOrder {
	return &ConfirmOrder{
		orderRepo: orderRepo,
	}
}

func (c ConfirmOrder) Do(ctx context.Context, orderId int64) error {
	return c.orderRepo.Confirm(ctx, orderId)
}
