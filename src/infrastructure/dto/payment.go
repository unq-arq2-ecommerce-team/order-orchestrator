package dto

import (
	"fmt"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
	"strconv"
)

type OrderPaymentReq struct {
	Method model.Method `json:"method" binding:"required"`
}

// TODO: Upgrade validation
func (req OrderPaymentReq) Validate() error {
	if !model.ValidPaymentMethodType(req.Method.Type) {
		return fmt.Errorf("payment method.type invalid. Please try with one of these: %s", model.GetPaymentMethodTypes())
	}
	if _, err := strconv.ParseUint(req.Method.Details.Number, 10, 64); err != nil {
		return fmt.Errorf("number is not valid. Please try with format: 1234123412341234")
	}
	if _, err := strconv.ParseUint(req.Method.Details.Cvv, 10, 64); err != nil {
		return fmt.Errorf("cvv is not valid. Please try with format: 999")
	}
	return nil
}

func (req OrderPaymentReq) Map() *model.Payment {
	return &model.Payment{
		Method: req.Method,
	}
}
