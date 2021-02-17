// Print a new notification in windows system

package notification

import (
	"gopkg.in/toast.v1"
)

// Create create a new notification with presets
func Create(title, message string, actions []toast.Action, loop bool) toast.Notification {
	createdNotif := toast.Notification{
		AppID:   "Gotomate",
		Title:   title,
		Message: message,
		Icon:    "C:/Users/lhopi/go/src/gotomate/img/gotomate.png",
		Actions: actions,
		Loop:    loop,
	}
	return createdNotif
}

// Push push a notification in user's system return nil if pushed else nil
func Push(notification toast.Notification) error {
	err := notification.Push()
	if err != nil {
		return nil
	}
	return err
}
