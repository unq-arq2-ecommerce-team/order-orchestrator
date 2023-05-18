package dto

import (
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
	"time"
)

type OrderDTO struct {
	CustomerId      int64         `json:"customerId"`
	ProductId       int64         `json:"productId"`
	DeliveryDate    time.Time     `json:"deliveryDate" time_format:"2006-01-02T15:04:05.000Z"`
	DeliveryAddress model.Address `json:"deliveryAddress"`
}

func NewOrderDTO(order model.Order) *OrderDTO {
	return &OrderDTO{
		CustomerId:      order.CustomerId,
		ProductId:       order.Product.Id,
		DeliveryDate:    order.DeliveryDate,
		DeliveryAddress: order.DeliveryAddress,
	}
}

type IdRes struct {
	Id int64 `json:"id"`
}

func NewIdRes(id int64) *IdResponse {
	return &IdResponse{Id: id}
}
