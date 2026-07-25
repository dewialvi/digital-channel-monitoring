package service

import (
	"fmt"
	"time"
)

type NotificationService struct{}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (s *NotificationService) SendNotification(
	user string,
	message string,
) {

	go func() {
		time.Sleep(2 * time.Second)

		fmt.Printf(
			"[NOTIFICATION] To: %s | Message: %s\n",
			user,
			message,
		)
	}()
}