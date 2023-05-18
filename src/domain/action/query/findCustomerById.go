package query

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
)

type FindCustomerById struct {
	customerRepo model.CustomerRepository
}

func NewFindCustomerById(customerRepo model.CustomerRepository) *FindCustomerById {
	return &FindCustomerById{
		customerRepo: customerRepo,
	}
}

func (q FindCustomerById) Do(ctx context.Context, id int64) (*model.Customer, error) {
	return q.customerRepo.FindById(ctx, id)
}
