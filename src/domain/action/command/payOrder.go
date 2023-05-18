package command

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
)

type PayOrder struct {
	paymentRepository model.PaymentRepository
}

func NewPayOrder(paymentRepository model.PaymentRepository) *PayOrder {
	return &PayOrder{
		paymentRepository: paymentRepository,
	}
}

func (c PayOrder) Do(ctx context.Context, payment *model.Payment) error {
	return c.paymentRepository.MakePayment(ctx, payment)
}
