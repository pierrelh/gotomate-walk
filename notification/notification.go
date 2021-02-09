// Print a new notification in windows system

package notification

import (
	"gopkg.in/toast.v1"
)

// CreateNotification create a new notification with presets
func CreateNotification(title string, message string, actions []toast.Action, loop bool) toast.Notification {
	createdNotif := toast.Notification{
		AppID:   "Gotomate",
		Title:   title,
		Message: message,
		Icon:    "C:/Users/lhopi/go/src/piproto/notification/favicon.ico",
		Actions: actions,
		Loop:    loop,
	}
	return createdNotif
}

// PushNotification push a notification in user's system return nil if pushed else nil
func PushNotification(notification toast.Notification) error {
	err := notification.Push()
	if err != nil {
		return nil
	}
	return err
}
