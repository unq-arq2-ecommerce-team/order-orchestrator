package model

import "github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/util"

type Address struct {
	Street      string `json:"street" binding:"required"`
	City        string `json:"city" binding:"required"`
	State       string `json:"state" binding:"required"`
	Country     string `json:"country" binding:"required"`
	Observation string `json:"observation"`
}

func (a Address) String() string {
	return util.ParseStruct("Address", a)
}
