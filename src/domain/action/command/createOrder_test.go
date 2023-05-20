package command

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/mock"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
	"testing"
	"time"
)

func Test_GivenCreateOrderCmdAndOrder_WhenDo_ThenReturnAnOrderIdAndNoError(t *testing.T) {
	createOrderCmd, mocks := setUpCreateOrderTest(t)
	ctx := context.Background()

	order := anyOrder()

	expectedOrderId := int64(15)
	mocks.OrderRepo.EXPECT().Create(ctx, order).Return(expectedOrderId, nil)

	orderIdCreated, err := createOrderCmd.Do(ctx, order)

	assert.Equal(t, expectedOrderId, orderIdCreated)
	assert.NoError(t, err)
}

func Test_GivenCreateOrderCmdWhichReturnErrorAndOrder_WhenDo_ThenReturnThatError(t *testing.T) {
	createOrderCmd, mocks := setUpCreateOrderTest(t)
	ctx := context.Background()

	order := anyOrder()

	errExpected := fmt.Errorf("whatever error")
	mocks.OrderRepo.EXPECT().Create(ctx, order).Return(int64(0), errExpected)

	orderIdCreated, err := createOrderCmd.Do(ctx, order)

	assert.Zero(t, orderIdCreated)
	assert.ErrorIs(t, err, errExpected)
}

func anyOrder() model.Order {
	return model.NewOrder(321, 5559, time.Now(), model.Address{})
}

func setUpCreateOrderTest(t *testing.T) (*CreateOrder, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewCreateOrder(mocks.OrderRepo), mocks
}
