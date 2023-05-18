package model

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/util"
)

type Customer struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

func (c *Customer) String() string {
	return util.ParseStruct("Customer", c)
}

//go:generate mockgen -destination=../mock/customerRepository.go -package=mock -source=customer.go
type CustomerRepository interface {
	FindById(ctx context.Context, id int64) (*Customer, error)
}
