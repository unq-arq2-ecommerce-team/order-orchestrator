package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Payment_String(t *testing.T) {
	p1 := Payment{}
	p2 := Payment{
		CustomerId: "2",
		OrderId:    "4",
		Amount:     105.95,
		Currency:   CurrencyARS,
		Method: Method{
			Type: "creditcard",
			Details: Details{
				Number:         "1234123412341234",
				Cvv:            "987",
				ExpirationDate: "12/2029",
				HolderName:     "Lalo Landa",
			},
		},
	}
	assert.Equal(t, `[Payment]{"customer_id":"","order_id":"","amount":0,"currency":"","method":{"type":"","details":{"number":"","expiration_date":"","cvv":"","holder_name":""}}}`, p1.String())
	assert.Equal(t, `[Payment]{"customer_id":"2","order_id":"4","amount":105.95,"currency":"ARS","method":{"type":"creditcard","details":{"number":"1234123412341234","expiration_date":"12/2029","cvv":"987","holder_name":"Lalo Landa"}}}`, p2.String())
}

func Test_Payment_Fill(t *testing.T) {
	method := Method{
		Type: "creditcard",
		Details: Details{
			Number:         "1234123412341234",
			Cvv:            "987",
			ExpirationDate: "12/2029",
			HolderName:     "Lalo Landa",
		},
	}
	payment := &Payment{
		Method:     method,
		CustomerId: "99999999",
		OrderId:    "99999999",
		Amount:     1,
		Currency:   "sarasa",
	}
	priceAmount := 105.95
	order := &Order{Id: 3, CustomerId: 87, Product: Product{Price: priceAmount}}

	payment.Fill(order)

	assert.Equal(t, method, payment.Method)
	assert.Equal(t, "3", payment.OrderId)
	assert.Equal(t, "87", payment.CustomerId)
	assert.Equal(t, priceAmount, payment.Amount)
	assert.Equal(t, CurrencyARS, payment.Currency)
}
