package command

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/mock"
	"github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
	"golang.org/x/net/context"
	"testing"
)

func Test_GivenSendNotificationCmdAndNotification_WhenDo_ThenReturnNoError(t *testing.T) {
	sendNotificationCmd, mocks := setUpSendNotificationTest(t)
	ctx := context.Background()

	notification := anyNotification()
	mocks.NotificationRepo.EXPECT().Send(ctx, notification).Return(nil)

	err := sendNotificationCmd.Do(ctx, notification)

	assert.NoError(t, err)
}

func Test_GivenSendNotificationCmdWhichReturnErrorAndNotification_WhenDo_ThenReturnThatError(t *testing.T) {
	sendNotificationCmd, mocks := setUpSendNotificationTest(t)
	ctx := context.Background()

	notification := anyNotification()
	errExpected := fmt.Errorf("some error in notification")
	mocks.NotificationRepo.EXPECT().Send(ctx, notification).Return(errExpected)

	err := sendNotificationCmd.Do(ctx, notification)

	assert.ErrorIs(t, err, errExpected)
}

func anyNotification() model.Notification {
	return model.Notification{}
}

func setUpSendNotificationTest(t *testing.T) (*SendNotification, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewSendNotification(mocks.NotificationRepo), mocks
}
