package exception

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CustomerNotFoundError(t *testing.T) {
	e := CustomerNotFound{
		Id: 1,
	}
	e2 := CustomerNotFound{
		Id: 2,
	}
	assert.Equal(t, `customer with id 1 not found`, e.Error())
	assert.Equal(t, `customer with id 2 not found`, e2.Error())
}
