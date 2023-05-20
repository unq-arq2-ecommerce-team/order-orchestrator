package usecase

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/action/command"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/mock"
	"golang.org/x/net/context"
	"testing"
)

func Test_GivenDeliveredOrderUseCaseAndOrderId_WhenDo_ThenReturnNoError(t *testing.T) {
	deliveredOrder, mocks := setUpDeliveredOrderTest(t)
	ctx := context.Background()

	orderId := int64(19)
	mocks.OrderRepo.EXPECT().Delivered(ctx, orderId).Return(nil)

	err := deliveredOrder.Do(ctx, orderId)

	assert.NoError(t, err)
}

func Test_GivenDeliveredOrderUseCaseWhichReturnErrorAndOrderId_WhenDo_ThenReturnThatError(t *testing.T) {
	deliveredOrder, mocks := setUpDeliveredOrderTest(t)
	ctx := context.Background()

	orderId := int64(79878)
	errExpected := fmt.Errorf("some error from delivered")
	mocks.OrderRepo.EXPECT().Delivered(ctx, orderId).Return(errExpected)

	err := deliveredOrder.Do(ctx, orderId)

	assert.ErrorIs(t, err, errExpected)
}

func setUpDeliveredOrderTest(t *testing.T) (*DeliveredOrder, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	deliveredOrderCmd := *command.NewDeliveredOrder(mocks.OrderRepo)
	return NewDeliveredOrder(mocks.Logger, deliveredOrderCmd), mocks
}
