package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Notification_String(t *testing.T) {
	n1 := Notification{
		Channel: "email",
		Event: Event{
			Name:   "event1",
			Detail: "detail1",
		},
		Recipient: Recipient{
			Type: "type",
			Id:   2,
		},
	}
	n2 := Notification{
		Channel: "sms",
		Event: Event{
			Name:   "event2",
			Detail: "detail2",
		},
		Recipient: Recipient{
			Type: "type2",
			Id:   65,
		},
	}
	assert.Equal(t, `[Notification]{"channel":"email","event":{"name":"event1","detail":"detail1"},"recipient":{"type":"type","id":2}}`, n1.String())
	assert.Equal(t, `[Notification]{"channel":"sms","event":{"name":"event2","detail":"detail2"},"recipient":{"type":"type2","id":65}}`, n2.String())
}

func Test_NewEmailNotificationOrderPayedForCustomer(t *testing.T) {
	recipientType, userId := CustomerRecipientType, int64(99)
	orderId, productName, productPrice := int64(123), "Yerba Matex", float64(99.99)
	order := anyOrderWithIdAndProductNameAndPrice(orderId, productPrice, productName)
	expectedDetail := "Yerba Matex - $99.99 - Numero de orden: #123"

	notification := NewEmailNotificationOrderPayed(recipientType, userId, order)
	expectedNotification := Notification{
		Channel: channelEmail,
		Event: Event{
			Name:   eventPaymentOk,
			Detail: expectedDetail,
		},
		Recipient: Recipient{
			Type: recipientType,
			Id:   userId,
		},
	}

	assert.Equal(t, expectedNotification, notification)
}

func Test_NewEmailNotificationOrderPayedForNoCustomer(t *testing.T) {
	recipientType, userId := SellerRecipientType, int64(99)
	orderId, productName, productPrice := int64(123), "Yerba Matex", float64(99.99)
	order := anyOrderWithIdAndProductNameAndPrice(orderId, productPrice, productName)
	expectedDetail := "Yerba Matex - $99.99"

	notification := NewEmailNotificationOrderPayed(recipientType, userId, order)
	expectedNotification := Notification{
		Channel: channelEmail,
		Event: Event{
			Name:   eventPaymentOk,
			Detail: expectedDetail,
		},
		Recipient: Recipient{
			Type: recipientType,
			Id:   userId,
		},
	}

	assert.Equal(t, expectedNotification, notification)
}

func anyOrderWithIdAndProductNameAndPrice(id int64, productPrice float64, productName string) Order {
	return Order{
		Id: id,
		Product: Product{
			Name:  productName,
			Price: productPrice,
		},
	}
}
