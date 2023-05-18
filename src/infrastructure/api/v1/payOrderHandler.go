package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model/exception"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/usecase"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/infrastructure/dto"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/infrastructure/logger"
	"net/http"
)

// PayOrderHandler
// @Summary      Endpoint order payment
// @Description  pay an order
// @Param orderId path int true "Order ID" minimum(1)
// @Param Order body dto.OrderPaymentReq true "It is a order payment request."
// @Tags         Order
// @Produce json
// @Success 204
// @Failure 400 {object} dto.ErrorMessage
// @Failure 404 {object} dto.ErrorMessage
// @Failure 406 {object} dto.ErrorMessage
// @Router       /api/v1/order/{orderId}/pay [post]
func PayOrderHandler(log model.Logger, payOrder *usecase.PayOrder) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parsePathParamPositiveIntId(c, "orderId")
		if err != nil {
			log.WithFields(logger.Fields{"error": err}).Error("invalid path param")
			writeJsonErrorMessageWithNoDesc(c, http.StatusBadRequest, err)
			return
		}
		var req dto.OrderPaymentReq
		if err := c.BindJSON(&req); err != nil {
			writeJsonErrorMessageInDescAndMessage(c, http.StatusBadRequest, "invalid json body order payment req", err)
			return
		}
		if err := req.Validate(); err != nil {
			writeJsonErrorMessageInDescAndMessage(c, http.StatusBadRequest, "invalid json body order payment req", err)
			return
		}
		if err := payOrder.Do(c.Request.Context(), id, req.Map()); err != nil {
			switch err.(type) {
			case exception.OrderNotFound:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotFound, err)
			case exception.OrderWasPaid:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotAcceptable, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when create order", err)
			}
			return
		}

		c.Status(http.StatusNoContent)
	}
}
