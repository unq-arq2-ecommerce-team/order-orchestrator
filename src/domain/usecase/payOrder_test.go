package usecase

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/action/command"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/mock"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
	"testing"
)

func Test_GivenPayOrderUseCaseAndOrderIdAndPayment_WhenDo_ThenReturnNoError(t *testing.T) {
	payOrderUseCase, mocks := setUpPayOrderTest(t)
	ctx := context.Background()

	orderId := int64(32)
	payment := anyPayment()

	order := anyOrderWithIdAndState(orderId, model.PendingOrderState)
	paymentForRepo := *payment
	paymentForRepo.Fill(order)
	mocks.OrderRepo.EXPECT().FindById(ctx, orderId).Return(order, nil)
	mocks.PaymentRepo.EXPECT().MakePayment(ctx, &paymentForRepo).Return(nil)
	mocks.OrderRepo.EXPECT().Confirm(ctx, order.Id).Return(nil)
	mocks.NotificationRepo.EXPECT().Send(ctx, gomock.Any()).Return(nil).Times(2)

	err := payOrderUseCase.Do(ctx, orderId, payment)

	assert.NoError(t, err)
}

func Test_GivenPayOrderUseCaseAndOrderIdAndPaymentAndNotificationError_WhenDo_ThenReturnNoError(t *testing.T) {
	payOrderUseCase, mocks := setUpPayOrderTest(t)
	ctx := context.Background()

	orderId := int64(32)
	payment := anyPayment()

	order := anyOrderWithIdAndState(orderId, model.PendingOrderState)
	paymentForRepo := *payment
	paymentForRepo.Fill(order)
	mocks.OrderRepo.EXPECT().FindById(ctx, orderId).Return(order, nil)
	mocks.PaymentRepo.EXPECT().MakePayment(ctx, &paymentForRepo).Return(nil)
	mocks.OrderRepo.EXPECT().Confirm(ctx, order.Id).Return(nil)
	mocks.NotificationRepo.EXPECT().Send(ctx, gomock.Any()).Return(fmt.Errorf("notification repo error")).Times(2)

	err := payOrderUseCase.Do(ctx, orderId, payment)

	assert.NoError(t, err)
}

func Test_GivenPayOrderUseCaseAndOrderIdAndPaymentAndOrderRepoConfirmErr_WhenDo_ThenReturnThatErr(t *testing.T) {
	payOrderUseCase, mocks := setUpPayOrderTest(t)
	ctx := context.Background()

	orderId := int64(32)
	payment := anyPayment()

	order := anyOrderWithIdAndState(orderId, model.PendingOrderState)
	paymentForRepo := *payment
	paymentForRepo.Fill(order)

	expectedErr := fmt.Errorf("some error in confirm in order repo")
	mocks.OrderRepo.EXPECT().FindById(ctx, orderId).Return(order, nil)
	mocks.PaymentRepo.EXPECT().MakePayment(ctx, &paymentForRepo).Return(nil)
	mocks.OrderRepo.EXPECT().Confirm(ctx, order.Id).Return(expectedErr)

	err := payOrderUseCase.Do(ctx, orderId, payment)

	assert.ErrorIs(t, err, expectedErr)
}

func Test_GivenPayOrderUseCaseAndOrderIdAndPaymentAndPaymentRepoErr_WhenDo_ThenReturnThatErr(t *testing.T) {
	payOrderUseCase, mocks := setUpPayOrderTest(t)
	ctx := context.Background()

	orderId := int64(32)
	payment := anyPayment()

	order := anyOrderWithIdAndState(orderId, model.PendingOrderState)
	paymentForRepo := *payment
	paymentForRepo.Fill(order)

	expectedErr := fmt.Errorf("some error in pament repo")
	mocks.OrderRepo.EXPECT().FindById(ctx, orderId).Return(order, nil)
	mocks.PaymentRepo.EXPECT().MakePayment(ctx, &paymentForRepo).Return(expectedErr)

	err := payOrderUseCase.Do(ctx, orderId, payment)

	assert.ErrorIs(t, err, expectedErr)
}

func Test_GivenPayOrderUseCaseAndOrderIdAndPaymentAndOrderRepoErr_WhenDo_ThenReturnThatErr(t *testing.T) {
	payOrderUseCase, mocks := setUpPayOrderTest(t)
	ctx := context.Background()

	orderId := int64(32)
	payment := anyPayment()

	order := anyOrderWithIdAndState(orderId, model.PendingOrderState)
	paymentForRepo := *payment
	paymentForRepo.Fill(order)

	expectedErr := fmt.Errorf("some error in order repo")
	mocks.OrderRepo.EXPECT().FindById(ctx, orderId).Return(order, expectedErr)

	err := payOrderUseCase.Do(ctx, orderId, payment)

	assert.ErrorIs(t, err, expectedErr)
}

func Test_GivenPayOrderUseCaseAndOrderIdWhichIsNoPendingOrderStateAndPayment_WhenDo_ThenReturnNoErrorForIdempotency(t *testing.T) {
	payOrderUseCase, mocks := setUpPayOrderTest(t)
	ctx := context.Background()

	order1Id := int64(32)
	order2Id := int64(32)
	payment1 := anyPayment()
	payment2 := anyPayment()

	order1 := anyOrderWithIdAndState(order1Id, model.ConfirmOrderState)
	order2 := anyOrderWithIdAndState(order2Id, model.DeliveredOrderState)
	mocks.OrderRepo.EXPECT().FindById(ctx, order1Id).Return(order1, nil)
	mocks.OrderRepo.EXPECT().FindById(ctx, order2Id).Return(order2, nil)

	err1 := payOrderUseCase.Do(ctx, order1Id, payment1)
	err2 := payOrderUseCase.Do(ctx, order2Id, payment2)

	assert.NoError(t, err1)
	assert.NoError(t, err2)
}

func setUpPayOrderTest(t *testing.T) (*PayOrder, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	findOrderByIdQuery := *query.NewFindOrderById(mocks.OrderRepo)
	confirmOrderCmd := *command.NewConfirmOrder(mocks.OrderRepo)
	sendNotificationCmd := *command.NewSendNotification(mocks.NotificationRepo)
	payOrderCmd := *command.NewPayOrder(mocks.PaymentRepo)
	return NewPayOrder(mocks.Logger, findOrderByIdQuery, payOrderCmd, confirmOrderCmd, sendNotificationCmd), mocks
}

func anyOrderWithIdAndState(id int64, state string) *model.Order {
	return &model.Order{Id: id, State: state, CustomerId: 123, Product: model.Product{Name: "Yerba", Price: 99.99, SellerId: 32}}
}

func anyPayment() *model.Payment {
	return &model.Payment{}
}
