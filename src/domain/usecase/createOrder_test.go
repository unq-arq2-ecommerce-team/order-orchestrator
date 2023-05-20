package usecase

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/action/command"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/mock"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
	"testing"
	"time"
)

func Test_GivenCreateOrderUseCaseAndOrder_WhenDo_ThenReturnAnOrderIdAndNoError(t *testing.T) {
	createOrderUseCase, mocks := setUpCreateOrderTest(t)
	ctx := context.Background()
	customer := anyCustomer()
	order := anyOrderWithCustomerId(customer.Id)

	expectedOrderId := int64(32)
	mocks.CustomerRepo.EXPECT().FindById(ctx, customer.Id).Return(customer, nil)
	mocks.OrderRepo.EXPECT().Create(ctx, order).Return(expectedOrderId, nil)

	orderId, err := createOrderUseCase.Do(ctx, order)

	assert.Equal(t, expectedOrderId, orderId)
	assert.NoError(t, err)
}

func Test_GivenCreateOrderUseCaseAndOrderAndCustomerRepoErr_WhenDo_ThenReturnThatCustomerRepoErrAndNoOrderId(t *testing.T) {
	createOrderUseCase, mocks := setUpCreateOrderTest(t)
	ctx := context.Background()
	customer := anyCustomer()
	order := anyOrderWithCustomerId(customer.Id)

	expectedErr := fmt.Errorf("some error found in customer repo")
	mocks.CustomerRepo.EXPECT().FindById(ctx, customer.Id).Return(nil, expectedErr)

	orderId, err := createOrderUseCase.Do(ctx, order)

	assert.Zero(t, orderId)
	assert.ErrorIs(t, err, expectedErr)
}

func Test_GivenCreateOrderUseCaseAndOrderAndOrderRepoErr_WhenDo_ThenReturnThatOrderRepoErrAndNoOrderId(t *testing.T) {
	createOrderUseCase, mocks := setUpCreateOrderTest(t)
	ctx := context.Background()
	customer := anyCustomer()
	order := anyOrderWithCustomerId(customer.Id)

	expectedErr := fmt.Errorf("some error found in order repo")
	mocks.CustomerRepo.EXPECT().FindById(ctx, customer.Id).Return(customer, nil)
	mocks.OrderRepo.EXPECT().Create(ctx, order).Return(int64(0), expectedErr)

	orderId, err := createOrderUseCase.Do(ctx, order)

	assert.Zero(t, orderId)
	assert.ErrorIs(t, err, expectedErr)
}

func setUpCreateOrderTest(t *testing.T) (*CreateOrder, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	findCustomerByIdQuery := *query.NewFindCustomerById(mocks.CustomerRepo)
	createOrderCmd := *command.NewCreateOrder(mocks.OrderRepo)
	return NewCreateOrder(mocks.Logger, findCustomerByIdQuery, createOrderCmd), mocks
}

func anyOrderWithCustomerId(customerId int64) model.Order {
	return model.NewOrder(customerId, 5559, time.Now(), model.Address{})
}

func anyCustomer() *model.Customer {
	return &model.Customer{Id: 32, Email: "sarasa@mail.com"}
}
