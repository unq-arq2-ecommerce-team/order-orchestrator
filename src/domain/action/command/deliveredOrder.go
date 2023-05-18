package command

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
)

type DeliveredOrder struct {
	orderRepo model.OrderRepository
}

func NewDeliveredOrder(orderRepo model.OrderRepository) *DeliveredOrder {
	return &DeliveredOrder{
		orderRepo: orderRepo,
	}
}

func (c DeliveredOrder) Do(ctx context.Context, orderId int64) error {
	return c.orderRepo.Delivered(ctx, orderId)
}
