package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Order_New(t *testing.T) {
	customerId, productId := int64(334), int64(998)
	deliveryDate := time.Date(2023, 4, 13, 3, 0, 0, 0, time.UTC)
	deliveryAddress := Address{
		Street:      "Fake street 123",
		City:        "La Plata",
		State:       "Buenos Aires",
		Country:     "Argentina",
		Observation: "asd",
	}
	order := NewOrder(customerId, productId, deliveryDate, deliveryAddress)

	orderExpected := Order{
		CustomerId:      customerId,
		Product:         Product{Id: productId},
		DeliveryAddress: deliveryAddress,
		DeliveryDate:    deliveryDate,
	}
	assert.Equal(t, orderExpected, order)
}

func Test_Order_WasPaid(t *testing.T) {
	pendingOrder := Order{
		State: PendingOrderState,
	}
	confirmedOrder := Order{
		State: ConfirmOrderState,
	}
	deliveredOrder := Order{
		State: DeliveredOrderState,
	}
	assert.False(t, pendingOrder.WasPaid())
	assert.True(t, confirmedOrder.WasPaid())
	assert.True(t, deliveredOrder.WasPaid())
}

func Test_Order_String(t *testing.T) {
	order1 := Order{
		CustomerId:      4,
		Product:         Product{Id: 43},
		State:           PendingOrderState,
		CreatedOn:       time.Date(2023, 4, 13, 3, 0, 0, 0, time.UTC),
		UpdatedOn:       time.Date(2023, 4, 13, 9, 0, 0, 0, time.UTC),
		DeliveryDate:    time.Date(2023, 4, 25, 16, 0, 0, 0, time.UTC),
		DeliveryAddress: Address{},
	}

	order2 := Order{
		Id:           5,
		CustomerId:   2,
		CreatedOn:    time.Date(2023, 8, 23, 3, 0, 0, 0, time.UTC),
		UpdatedOn:    time.Date(2024, 8, 23, 9, 0, 0, 0, time.UTC),
		DeliveryDate: time.Date(2023, 10, 25, 15, 0, 0, 0, time.UTC),
		State:        ConfirmOrderState,
		DeliveryAddress: Address{
			Street:      "Fake street 123",
			City:        "La Plata",
			State:       "Buenos Aires",
			Country:     "Argentina",
			Observation: "asd",
		},
	}

	order3 := Order{
		CustomerId:   6,
		Product:      Product{Id: 17, SellerId: 2, Name: "fafafa", Price: 10.5},
		State:        DeliveredOrderState,
		CreatedOn:    time.Date(2022, 4, 13, 3, 0, 0, 0, time.UTC),
		UpdatedOn:    time.Date(2022, 4, 13, 9, 0, 0, 0, time.UTC),
		DeliveryDate: time.Date(2022, 4, 25, 16, 0, 0, 0, time.UTC),
		DeliveryAddress: Address{
			Street:      "Fake street 123",
			City:        "La Plata",
			State:       "Buenos Aires",
			Country:     "Argentina",
			Observation: "asd",
		},
	}
	assert.Equal(t, `[Order]{"id":0,"customerId":4,"product":{"productId":43,"sellerId":0,"name":"","price":0},"state":"PENDING","createdOn":"2023-04-13T03:00:00Z","updatedOn":"2023-04-13T09:00:00Z","deliveryDate":"2023-04-25T16:00:00Z","deliveryAddress":{"street":"","city":"","state":"","country":"","observation":""}}`, order1.String())
	assert.Equal(t, `[Order]{"id":5,"customerId":2,"product":{"productId":0,"sellerId":0,"name":"","price":0},"state":"CONFIRMED","createdOn":"2023-08-23T03:00:00Z","updatedOn":"2024-08-23T09:00:00Z","deliveryDate":"2023-10-25T15:00:00Z","deliveryAddress":{"street":"Fake street 123","city":"La Plata","state":"Buenos Aires","country":"Argentina","observation":"asd"}}`, order2.String())
	assert.Equal(t, `[Order]{"id":0,"customerId":6,"product":{"productId":17,"sellerId":2,"name":"fafafa","price":10.5},"state":"DELIVERED","createdOn":"2022-04-13T03:00:00Z","updatedOn":"2022-04-13T09:00:00Z","deliveryDate":"2022-04-25T16:00:00Z","deliveryAddress":{"street":"Fake street 123","city":"La Plata","state":"Buenos Aires","country":"Argentina","observation":"asd"}}`, order3.String())

}
