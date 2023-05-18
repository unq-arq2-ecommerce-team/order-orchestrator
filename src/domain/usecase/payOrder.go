package usecase

import (
	"context"
	"fmt"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/action/command"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model/exception"
)

type PayOrder struct {
	baseLogger          model.Logger
	findOrderByIdQuery  query.FindOrderById
	confirmOrderCmd     command.ConfirmOrder
	sendNotificationCmd command.SendNotification
	payOrderCmd         command.PayOrder
}

func NewPayOrder(baseLogger model.Logger, findOrderByIdQuery query.FindOrderById, payOrderCmd command.PayOrder, confirmOrderCmd command.ConfirmOrder, sendNotificationCmd command.SendNotification) *PayOrder {
	return &PayOrder{
		baseLogger:          baseLogger.WithFields(model.LoggerFields{"useCase": "PayOrder"}),
		findOrderByIdQuery:  findOrderByIdQuery,
		payOrderCmd:         payOrderCmd,
		confirmOrderCmd:     confirmOrderCmd,
		sendNotificationCmd: sendNotificationCmd,
	}
}

func (u *PayOrder) Do(ctx context.Context, orderId int64, payment *model.Payment) error {
	log := u.baseLogger.WithFields(model.LoggerFields{"orderId": orderId, "payment": payment})
	log.Debug("paying order ...")
	order, err := u.findOrderByIdQuery.Do(ctx, orderId)
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Errorf("error when find order by id %v", orderId)
		return err
	}
	log = log.WithFields(model.LoggerFields{"order": order})

	if order.WasPaid() {
		log.Errorf("cannot pay order because not apply, orderState: %s", order.State)
		return exception.OrderWasPaid{Id: orderId}
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

	u.sendNotificationsOfPurchasedOrder(ctx, log, order)
	log.Infof("successfully paid order")
	return nil
}

func (u *PayOrder) sendNotificationsOfPurchasedOrder(ctx context.Context, log model.Logger, order *model.Order) {
	if err := u.sendNotificationCmd.Do(ctx, model.NewNotificationOrderPayed(model.CustomerRecipientType, order.CustomerId, fmt.Sprintf("%s - Numero de orden: #%v", order.Product.Name, order.Id))); err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("error when notify order purchased to customer")
	}
	if err := u.sendNotificationCmd.Do(ctx, model.NewNotificationOrderPayed(model.SellerRecipientType, order.Product.SellerId, fmt.Sprintf("%s", order.Product.Name))); err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("error when notify order purchased to seller")
	}
}
