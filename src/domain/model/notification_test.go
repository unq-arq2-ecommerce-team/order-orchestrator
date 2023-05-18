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

func Test_NewEmailNotificationOrderPayed(t *testing.T) {
	recientType, userId, customDetail := "recientType", int64(21), "some detail bla bla"
	notification := NewEmailNotificationOrderPayed(recientType, userId, customDetail)
	expectedNotification := Notification{
		Channel: channelEmail,
		Event: Event{
			Name:   eventPaymentOk,
			Detail: customDetail,
		},
		Recipient: Recipient{
			Type: recientType,
			Id:   userId,
		},
	}

	assert.Equal(t, expectedNotification, notification)
}
