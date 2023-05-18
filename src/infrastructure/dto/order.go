package dto

import (
	"fmt"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
	"time"
)

type OrderCreateReq struct {
	CustomerId      int64         `json:"customerId" binding:"required,min=1"`
	ProductId       int64         `json:"productId" binding:"required,min=1"`
	DeliveryDate    time.Time     `json:"deliveryDate" time_format:"2006-01-02T15:04:05.000Z" example:"2090-04-20T15:04:05.000Z" binding:"required"`
	DeliveryAddress model.Address `json:"deliveryAddress" binding:"required"`
}

func (req *OrderCreateReq) Validate() error {
	timeNow := time.Now()
	if req.DeliveryDate.Before(timeNow) {
		return fmt.Errorf("invalid delivery date because is before to %s", timeNow)
	}
	return nil
}

// TODO: Add fields for payment
type OrderPaymentReq struct {
}
