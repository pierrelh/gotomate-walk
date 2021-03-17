// Print a new notification in windows system

package notification

import (
	"fmt"

	"github.com/lxn/walk"
)

// Create create a new notification with presets & push it
func Create(title, message string, finished chan bool) {
	fmt.Println("FIBER INFO: Creating notification ...")

	notTitle := "Default Title"
	notMsg := "Default Message"
	if title != "" {
		notTitle = title
	}
	if message != "" {
		notMsg = message
	}

	mw, err := walk.NewMainWindow()
	if err != nil {
		fmt.Println("FIBER ERROR: ", err)
	}

	icon, err := walk.Resources.Icon("/img/icon.ico")
	if err != nil {
		fmt.Println("FIBER ERROR: ", err)
	}

	ni, err := walk.NewNotifyIcon(mw)
	if err != nil {
		fmt.Println("FIBER ERROR: ", err)
	}

	if err := ni.SetIcon(icon); err != nil {
		fmt.Println("FIBER ERROR: ", err)
	}

	if err := ni.SetVisible(true); err != nil {
		fmt.Println("FIBER ERROR: ", err)
	}

	if err := ni.ShowInfo(notTitle, notMsg); err != nil {
		fmt.Println("FIBER ERROR: ", err)
	}

	finished <- true

	mw.Run()
}
