// Print a new notification in windows system

package notification

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"

	"github.com/lxn/walk"
)

// Create create a new notification with presets & push it
func Create(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Creating notification ...")

	var title string
	if titleIsVar := instructionData.FieldByName("TitleIsVar").Interface().(bool); titleIsVar {
		titleVarName := instructionData.FieldByName("TitleVarName").Interface().(string)
		if val := variable.SearchVariable(titleVarName).Value; val != nil {
			title = val.(string)
		} else {
			finished <- true
			return -1
		}
	} else {
		title = instructionData.FieldByName("Title").Interface().(string)
	}

	var msg string
	if msgIsVar := instructionData.FieldByName("MessageIsVar").Interface().(bool); msgIsVar {
		messageVarName := instructionData.FieldByName("MessageVarName").Interface().(string)
		if val := variable.SearchVariable(messageVarName).Value; val != nil {
			msg = val.(string)
		} else {
			finished <- true
			return -1
		}
	} else {
		msg = instructionData.FieldByName("Message").Interface().(string)
	}

	notTitle := "Default Title"
	notMsg := "Default Message"
	if title != "" {
		notTitle = title
	}
	if msg != "" {
		notMsg = msg
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
	return -1
}
