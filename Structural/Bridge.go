package Structural

import (
	"errors"
	"github.com/labstack/gommon/log"
)

// It separates the abstraction from its implementation and implements the abstraction in terms of implementor.

// In this the implementation part would be handled by the notification sender and the Notification will not concern itself with the send procedure. It delegates the same to the sender

type Notification interface {
	Send()
	SetNotificationSender(sender NotificationSender) error
}

type PushNotification struct {
	NotificationSender NotificationSender
}

func (notification *PushNotification) Send() {
	notification.NotificationSender.SendNotification()
}

func (notification *PushNotification) SetNotificationSender(notifSender NotificationSender) error {
	if notifSender != nil {
		notification.NotificationSender = notifSender
		return nil
	} else {
		return errors.New("Notification Sender not found")
	}
}

type TextNotification struct {
	NotificationSender NotificationSender
}

func (notification *TextNotification) Send() {
	notification.NotificationSender.SendNotification()
}

func (notification *TextNotification) SetNotificationSender(notifSender NotificationSender) error {
	if notifSender != nil {
		notification.NotificationSender = notifSender
		return nil
	} else {
		return errors.New("Notification Sender not found")
	}
}

type NotificationSender interface {
	SendNotification()
}

type PushNotificationSender struct {
}

func (notificationSender *PushNotificationSender) SendNotification() {
	log.Info("Push Notification!!")
}

type TextNotificationSender struct {
}

func (notificationSender *TextNotificationSender) SendNotification() {
	log.Info("Text Notification!!")
}

func PublishNotification(notiType string) {

	var notificationSender NotificationSender

	var notification Notification
	if notiType == "Push" {
		notificationSender = new(PushNotificationSender)
		notification = new(PushNotification)
	} else if notiType == "Text" {
		notificationSender = new(TextNotificationSender)
		notification = new(TextNotification)
	}

	if &notificationSender != nil {
		err := notification.SetNotificationSender(notificationSender)
		if err == nil {
			notification.Send()
		}
	}

}
