package model

import (
	"context"
	"fmt"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/util"
)

const CurrencyARS = "ARS"

type (
	Payment struct {
		CustomerId string  `json:"customer_id"`
		OrderId    string  `json:"order_id"`
		Amount     float64 `json:"amount"`
		Currency   string  `json:"currency"`
		Method     Method  `json:"method"`
	}
	Method struct {
		Type    string  `json:"type" binding:"required,min=1,max=20"`
		Details Details `json:"details" binding:"required"`
	}
	// Details ExpirationDate format is "mm/yyyy"
	Details struct {
		Number         string `json:"number" binding:"required,len=16"`
		ExpirationDate string `json:"expiration_date" binding:"required,len=7"`
		Cvv            string `json:"cvv" binding:"required,len=3"`
		HolderName     string `json:"holder_name" binding:"required,min=2,max=60"`
	}
)

// Fill : update payment with order data (orderId, customerId and price in ARS)
func (p *Payment) Fill(order *Order) {
	p.OrderId = fmt.Sprintf("%v", order.Id)
	p.CustomerId = fmt.Sprintf("%v", order.CustomerId)
	p.Amount = order.Product.Price
	p.Currency = CurrencyARS
}

func (p *Payment) String() string {
	return util.ParseStruct("Payment", p)
}

//go:generate mockgen -destination=../mock/paymentRepository.go -package=mock -source=payment.go
type PaymentRepository interface {
	MakePayment(ctx context.Context, payment *Payment) error
}
