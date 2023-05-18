package model

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/util"
	"strings"
	"time"
)

const (
	PendingOrderState   = "PENDING"
	ConfirmOrderState   = "CONFIRMED"
	DeliveredOrderState = "DELIVERED"
)

type (
	Order struct {
		Id              int64     `json:"id"`
		CustomerId      int64     `json:"customerId"`
		Product         Product   `json:"product"`
		State           string    `json:"state"`
		CreatedOn       time.Time `json:"createdOn"`
		UpdatedOn       time.Time `json:"updatedOn"`
		DeliveryDate    time.Time `json:"deliveryDate"`
		DeliveryAddress Address   `json:"deliveryAddress"`
	}
	Product struct {
		Id       int64   `json:"productId"`
		SellerId int64   `json:"sellerId"`
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
	}
)

func NewOrder(customerId, productId int64, deliveryDate time.Time, deliveryAddress Address) Order {
	return Order{
		CustomerId:      customerId,
		Product:         Product{Id: productId},
		DeliveryDate:    deliveryDate,
		DeliveryAddress: deliveryAddress,
	}
}

func (o *Order) WasPaid() bool {
	return !strings.EqualFold(o.State, PendingOrderState)
}
func (o *Order) String() string {
	return util.ParseStruct("Order", o)
}

//go:generate mockgen -destination=../mock/orderRepository.go -package=mock -source=order.go
type OrderRepository interface {
	FindById(ctx context.Context, id int64) (*Order, error)
	Create(ctx context.Context, order Order) (int64, error)
	Confirm(ctx context.Context, id int64) error
	Delivered(ctx context.Context, id int64) error
}
