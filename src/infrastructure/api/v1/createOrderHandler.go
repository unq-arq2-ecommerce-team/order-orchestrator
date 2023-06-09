package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model/exception"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/usecase"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/infrastructure/dto"
	"net/http"
)

// CreateOrderHandler
// @Summary      Endpoint create order
// @Description  create order
// @Param Order body dto.OrderCreateReq true "It is a order creation request."
// @Tags         Order
// @Produce json
// @Success 200 {object} dto.IdResponse
// @Failure 400 {object} dto.ErrorMessage
// @Failure 404 {object} dto.ErrorMessage
// @Failure 406 {object} dto.ErrorMessage
// @Router       /api/v1/order [post]
func CreateOrderHandler(log model.Logger, createOrder *usecase.CreateOrder) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.OrderCreateReq
		if err := c.BindJSON(&req); err != nil {
			writeJsonErrorMessageInDescAndMessage(c, http.StatusBadRequest, "invalid json body order create req", err)
			return
		}
		if err := req.Validate(); err != nil {
			writeJsonErrorMessageWithNoDesc(c, http.StatusBadRequest, err)
			return
		}

		orderId, err := createOrder.Do(c.Request.Context(), model.NewOrder(req.CustomerId, req.ProductId, req.DeliveryDate, req.DeliveryAddress))
		if err != nil {
			switch err.(type) {
			case exception.CustomerNotFound, exception.ProductNotFound:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotFound, err)
			case exception.ProductWithNoStock:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotAcceptable, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when create order", err)
			}
			return
		}

		c.JSON(http.StatusOK, dto.NewIdResponse(orderId))
	}
}
