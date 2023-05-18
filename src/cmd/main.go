package main

import (
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/action/command"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/usecase"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/infrastructure/api"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/infrastructure/config"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/infrastructure/logger"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/infrastructure/repository/http"
)

func main() {
	conf := config.LoadConfig()
	baseLogger := logger.New(&logger.Config{
		ServiceName:     "order-orchestrator",
		EnvironmentName: conf.Environment,
		LogLevel:        conf.LogLevel,
		LogFormat:       logger.JsonFormat,
	})

	// repositories
	customerRepo := http.NewCustomerRepository(baseLogger, http.NewClient(), &conf.CustomerUrl)
	orderRepo := http.NewOrderRepository(baseLogger, http.NewClient(), &conf.OrderUrl)
	notificationRepo := http.NewNotificationRepository(baseLogger, http.NewClient(), &conf.NotificationUrl)
	paymentRepo := http.NewPaymentRepository(baseLogger, http.NewClient(), &conf.PaymentUrl)

	// queries
	findCustomerByIdQuery := query.NewFindCustomerById(customerRepo)
	findOrderByIdQuery := query.NewFindOrderById(orderRepo)

	// commands
	createOrderCmd := command.NewCreateOrder(orderRepo)
	confirmOrderCmd := command.NewConfirmOrder(orderRepo)
	deliveredOrderCmd := command.NewDeliveredOrder(orderRepo)
	sendNotificationCmd := command.NewSendNotification(notificationRepo)
	payOrderCmd := command.NewPayOrder(paymentRepo)

	// use cases
	createOrderUseCase := usecase.NewCreateOrder(baseLogger, *findCustomerByIdQuery, *createOrderCmd)
	payOrderUseCase := usecase.NewPayOrder(baseLogger, *findOrderByIdQuery, *payOrderCmd, *confirmOrderCmd, *sendNotificationCmd)
	deliveredOrderUseCase := usecase.NewDeliveredOrder(baseLogger, *deliveredOrderCmd)

	app := api.NewApplication(baseLogger, conf, &api.ApplicationUseCases{
		CreateOrderUseCase:    createOrderUseCase,
		PayOrderUseCase:       payOrderUseCase,
		DeliveredOrderUseCase: deliveredOrderUseCase,
	})
	baseLogger.Fatal(app.Run())
}
