package dto

import (
	"fmt"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
	"strconv"
)

type OrderPaymentReq struct {
	Method model.Method `json:"method" binding:"required"`
}

func (req OrderPaymentReq) Validate() error {
	if number, err := strconv.Atoi(req.Method.Details.Number); err != nil || number < 0 {
		return fmt.Errorf("number is not valid")
	}
	if cvv, err := strconv.Atoi(req.Method.Details.Cvv); err != nil || cvv < 0 {
		return fmt.Errorf("cvv is not valid")
	}
	return nil
}

func (req OrderPaymentReq) Map() *model.Payment {
	return &model.Payment{
		Method: req.Method,
	}
}
