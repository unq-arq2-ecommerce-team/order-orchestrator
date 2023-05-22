package usecase

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/action/command"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
)

type PayOrder struct {
	baseLogger          model.Logger
	findOrderByIdQuery  query.FindOrderById
	confirmOrderCmd     command.ConfirmOrder
	sendNotificationCmd command.SendNotification
	payOrderCmd         command.PayOrder
}

func NewPayOrder(
	baseLogger model.Logger, findOrderByIdQuery query.FindOrderById, payOrderCmd command.PayOrder,
	confirmOrderCmd command.ConfirmOrder, sendNotificationCmd command.SendNotification,
) *PayOrder {
	return &PayOrder{
		baseLogger:          baseLogger.WithFields(model.LoggerFields{"useCase": "PayOrder"}),
		findOrderByIdQuery:  findOrderByIdQuery,
		payOrderCmd:         payOrderCmd,
		confirmOrderCmd:     confirmOrderCmd,
		sendNotificationCmd: sendNotificationCmd,
	}
}

func (u *PayOrder) Do(ctx context.Context, orderId int64, payment *model.Payment) error {
	log := u.baseLogger.WithRequestId(ctx).WithFields(model.LoggerFields{"orderId": orderId, "payment": payment})
	log.Debug("paying order ...")
	order, err := u.findOrderByIdQuery.Do(ctx, orderId)
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Errorf("error when find order by id %v", orderId)
		return err
	}
	log = log.WithFields(model.LoggerFields{"order": order})

	// Idempotent
	if order.WasPaid() {
		log.Info("order with id %v state %s was already paid", order.Id, order.State)
		return nil
	}

	payment.Fill(order)
	log.Debugf("Payment filled: %s", payment)

	// Idempotent
	if err := u.payOrderCmd.Do(ctx, payment); err != nil {
		log.WithFields(model.LoggerFields{"error": err, "payment": payment}).Errorf("error when pay order")
		return err
	}

	// Idempotent
	if err := u.confirmOrderCmd.Do(ctx, order.Id); err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Errorf("error when confirm order by id %v", orderId)
		return err
	}

	// Async call
	go u.sendNotificationsOfPurchasedOrder(log, order)

	log.Infof("successfully paid order")
	return nil
}

func (u *PayOrder) sendNotificationsOfPurchasedOrder(log model.Logger, order *model.Order) {
	// create new ctx for no cancel from parent
	ctx := context.Background()
	if err := u.sendNotificationCmd.Do(ctx, model.NewEmailNotificationOrderPayed(model.CustomerRecipientType, order.CustomerId, *order)); err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("error when notify order purchased to customer")
	}
	if err := u.sendNotificationCmd.Do(ctx, model.NewEmailNotificationOrderPayed(model.SellerRecipientType, order.Product.SellerId, *order)); err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("error when notify order purchased to seller")
	}
}
