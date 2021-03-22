package packages

import (
	"fmt"
	algorithmic "gotomate/fiber/packages/Algorithmic"
	battery "gotomate/fiber/packages/Battery"
	clipboard "gotomate/fiber/packages/Clipboard"
	flow "gotomate/fiber/packages/Flow"
	input "gotomate/fiber/packages/Input"
	keyboard "gotomate/fiber/packages/Keyboard"
	log "gotomate/fiber/packages/Log"
	mouse "gotomate/fiber/packages/Mouse"
	notification "gotomate/fiber/packages/Notification"
	process "gotomate/fiber/packages/Process"
	screen "gotomate/fiber/packages/Screen"
	sleep "gotomate/fiber/packages/Sleep"
	systime "gotomate/fiber/packages/Systime"

	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
)

// Dialog Setting the requirement of a dialog window
type Dialog struct {
	Dialog        *walk.Dialog
	DialogContent declarative.Dialog
	DataBinder    *walk.DataBinder
	AcceptButton  *walk.PushButton
	CancelButton  *walk.PushButton
}

// CreateNewDialog Create the dialog for the function & add the needed template
func CreateNewDialog(packageName string, funcName string, databinder ...interface{}) (interface{}, *Dialog) {
	dialog := new(Dialog)
	source, children := PackageDecode(packageName, funcName)

	if len(databinder) != 0 {
		source = databinder[0]
	}

	dialog.DialogContent = declarative.Dialog{
		Icon:          "/img/icon.ico",
		Title:         funcName + " Settings",
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		AssignTo:      &dialog.Dialog,
		DefaultButton: &dialog.AcceptButton,
		CancelButton:  &dialog.CancelButton,
		DataBinder: declarative.DataBinder{
			AssignTo:       &dialog.DataBinder,
			Name:           funcName,
			DataSource:     source,
			ErrorPresenter: declarative.ToolTipErrorPresenter{},
		},
		MinSize: declarative.Size{
			Width:  300,
			Height: 300,
		},
		Layout: declarative.VBox{},
		Children: []declarative.Widget{
			declarative.Composite{
				Layout:   declarative.VBox{},
				Children: children,
			},
			declarative.Composite{
				Alignment: declarative.Alignment2D(walk.AlignHCenterVCenter),
				Layout:    declarative.HBox{},
				Children: []declarative.Widget{
					declarative.PushButton{
						AssignTo: &dialog.AcceptButton,
						Text:     "OK",
						Font:     declarative.Font{Family: "Roboto", PointSize: 9},
						OnClicked: func() {
							if err := dialog.DataBinder.Submit(); err != nil {
								fmt.Println("GOTOMATE ERROR: Unable to send the dialog's datas")
								return
							}

							dialog.Dialog.Accept()
						},
					},
					declarative.PushButton{
						AssignTo:  &dialog.CancelButton,
						Text:      "Cancel",
						Font:      declarative.Font{Family: "Roboto", PointSize: 9},
						OnClicked: func() { dialog.Dialog.Cancel() },
					},
				},
			},
		},
	}

	return source, dialog
}

// PackageDecode Getting the right databinder & the right template needed
func PackageDecode(packageName string, funcName string) (interface{}, []declarative.Widget) {
	switch packageName {
	case "Battery":
		return battery.Build(funcName)
	case "Clipboard":
		return clipboard.Build(funcName)
	case "Flow":
		return flow.Build(funcName)
	case "Algorithmic":
		return algorithmic.Build(funcName)
	case "Input":
		return input.Build(funcName)
	case "Keyboard":
		return keyboard.Build(funcName)
	case "Log":
		return log.Build(funcName)
	case "Mouse":
		return mouse.Build(funcName)
	case "Notification":
		return notification.Build(funcName)
	case "Process":
		return process.Build(funcName)
	case "Screen":
		return screen.Build(funcName)
	case "Sleep":
		return sleep.Build(funcName)
	case "Systime":
		return systime.Build(funcName)
	default:
		fmt.Println("GOTOMATE ERROR: Unable to find the dialog's package")
		return nil, nil
	}
}
