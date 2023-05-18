package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Customer_String(t *testing.T) {
	customer1 := Customer{
		Email: "pepegrillo@mail.com",
	}
	customer2 := Customer{
		Id:    2,
		Email: "sarasa@mail.com",
	}
	assert.Equal(t, `[Customer]{"id":0,"email":"pepegrillo@mail.com"}`, customer1.String())
	assert.Equal(t, `[Customer]{"id":2,"email":"sarasa@mail.com"}`, customer2.String())
}
