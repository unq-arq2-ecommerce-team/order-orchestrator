package exception

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_OrderNotFoundError(t *testing.T) {
	e1 := OrderNotFound{
		Id: 1,
	}
	e2 := OrderNotFound{
		Id: 2,
	}
	assert.Equal(t, `order with id 1 not found`, e1.Error())
	assert.Equal(t, `order with id 2 not found`, e2.Error())
}

func Test_OrderInvalidTransitionStateError(t *testing.T) {
	e1 := OrderInvalidTransitionState{
		Id: 1,
	}
	e2 := OrderInvalidTransitionState{
		Id: 2,
	}
	assert.Equal(t, `invalid transition state for order with id 1`, e1.Error())
	assert.Equal(t, `invalid transition state for order with id 2`, e2.Error())
}

func Test_OrderWasPaidError(t *testing.T) {
	e1 := OrderWasPaid{
		Id: 1,
	}
	e2 := OrderWasPaid{
		Id: 2,
	}
	assert.Equal(t, `order was paid for order with id 1`, e1.Error())
	assert.Equal(t, `order was paid for order with id 2`, e2.Error())
}
