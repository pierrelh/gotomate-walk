// Print a new notification in windows system

package notification

import (
	"fmt"

	"gopkg.in/toast.v1"
)

// Create create a new notification with presets & push it
func Create(title, message string, actions []toast.Action, finished chan bool) {
	fmt.Println("Notification initialization")
	createdNotif := toast.Notification{
		AppID:   "Gotomate",
		Title:   title,
		Message: message,
		Icon:    "C:/Users/lhopi/go/src/gotomate/img/gotomate.png",
		Actions: actions,
	}
	createdNotif.Push()
	finished <- true
}
