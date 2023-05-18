package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerDocs "github.com/unq-arq2-ecommerce-team/order-orchestrator/docs"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/usecase"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/infrastructure/api/middleware"
	v1 "github.com/unq-arq2-ecommerce-team/order-orchestrator/src/infrastructure/api/v1"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/infrastructure/config"
	"io"
	"net/http"
)

// Application
// @title order-orchestrator API
// @version 1.0
// @description api for tp arq2
// @contact.name API SUPPORT
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @BasePath /
// @query.collection.format multi
type Application interface {
	Run() error
}

type application struct {
	logger model.Logger
	config config.Config
	*ApplicationUseCases
}

type ApplicationUseCases struct {
	CreateOrderUseCase    *usecase.CreateOrder
	PayOrderUseCase       *usecase.PayOrder
	DeliveredOrderUseCase *usecase.DeliveredOrder
}

func NewApplication(l model.Logger, conf config.Config, applicationUseCases *ApplicationUseCases) Application {
	return &application{
		logger:              l,
		config:              conf,
		ApplicationUseCases: applicationUseCases,
	}
}

func (app *application) Run() error {
	swaggerDocs.SwaggerInfo.Host = fmt.Sprintf("localhost:%v", app.config.Port)

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	router := gin.Default()
	router.GET("/", HealthCheck)

	rv1 := router.Group("/api/v1")
	rv1.Use(middleware.TracingRequestId())
	rv1.POST("/order", v1.CreateOrderHandler(app.logger, app.CreateOrderUseCase))
	rv1.POST("/order/:orderId/pay", v1.PayOrderHandler(app.logger, app.PayOrderUseCase))
	rv1.POST("/order/:orderId/delivered", v1.DeliveredOrderHandler(app.logger, app.DeliveredOrderUseCase))

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.logger.Infof("running http server on port %d", app.config.Port)
	return router.Run(fmt.Sprintf(":%d", app.config.Port))
}

// HealthCheck
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags Health check
// @Accept */*
// @Produce json
// @Success 200 {object} HealthCheckRes
// @Router / [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, HealthCheckRes{Data: "Server is up and running"})
}

type HealthCheckRes struct {
	Data string `json:"data"`
}
