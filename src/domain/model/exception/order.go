package exception

import "fmt"

type OrderNotFound struct {
	Id int64
}

func (e OrderNotFound) Error() string {
	return fmt.Sprintf("order with id %v not found", e.Id)
}

type OrderInvalidTransitionState struct {
	Id int64
}

func (e OrderInvalidTransitionState) Error() string {
	return fmt.Sprintf("invalid transition state for order with id %v", e.Id)
}

type OrderWasPaid struct {
	Id int64
}

func (e OrderWasPaid) Error() string {
	return fmt.Sprintf("order was paid for order with id %v", e.Id)
}
