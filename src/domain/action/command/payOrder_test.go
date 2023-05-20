package command

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/mock"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
	"testing"
)

func Test_GivenPayOrderCmdAndAPayment_WhenDo_ThenReturnNoError(t *testing.T) {
	payOrderCmd, mocks := setUpPayOrderTest(t)
	ctx := context.Background()

	payment := anyPayment()
	mocks.PaymentRepo.EXPECT().MakePayment(ctx, payment).Return(nil)

	err := payOrderCmd.Do(ctx, payment)

	assert.NoError(t, err)
}

func Test_GivenPayOrderCmdWhichReturnErrorAndPayment_WhenDo_ThenReturnThatError(t *testing.T) {
	payOrderCmd, mocks := setUpPayOrderTest(t)
	ctx := context.Background()

	payment := anyPayment()
	errExpected := fmt.Errorf("some error")
	mocks.PaymentRepo.EXPECT().MakePayment(ctx, payment).Return(errExpected)

	err := payOrderCmd.Do(ctx, payment)

	assert.ErrorIs(t, err, errExpected)
}

func anyPayment() *model.Payment {
	return &model.Payment{}
}

func setUpPayOrderTest(t *testing.T) (*PayOrder, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewPayOrder(mocks.PaymentRepo), mocks
}
