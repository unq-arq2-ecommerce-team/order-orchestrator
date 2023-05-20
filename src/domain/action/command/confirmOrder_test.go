package command

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/mock"
	"testing"
)

func Test_GivenConfirmOrderCmdAndOrderId_WhenDo_ThenReturnNoError(t *testing.T) {
	confirmOrderCmd, mocks := setUpConfirmOrderTest(t)
	ctx := context.Background()

	orderId := int64(15)
	mocks.OrderRepo.EXPECT().Confirm(ctx, orderId).Return(nil)

	err := confirmOrderCmd.Do(ctx, orderId)

	assert.NoError(t, err)
}

func Test_GivenConfirmOrderCmdWhichReturnErrorAndOrderId_WhenDo_ThenReturnThatError(t *testing.T) {
	confirmOrderCmd, mocks := setUpConfirmOrderTest(t)
	ctx := context.Background()

	orderId := int64(15)
	errExpected := fmt.Errorf("some error")
	mocks.OrderRepo.EXPECT().Confirm(ctx, orderId).Return(errExpected)

	err := confirmOrderCmd.Do(ctx, orderId)

	assert.ErrorIs(t, err, errExpected)
}

func setUpConfirmOrderTest(t *testing.T) (*ConfirmOrder, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewConfirmOrder(mocks.OrderRepo), mocks
}
