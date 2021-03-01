package packages

import (
	"fmt"

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
func CreateNewDialog(funcName string, databinder ...interface{}) (interface{}, *Dialog) {
	dialog := new(Dialog)
	source, children := FillDialog(funcName)

	if len(databinder) != 0 {
		source = databinder[0]
	}

	dialog.DialogContent = declarative.Dialog{
		Icon:          "/icon.ico",
		Title:         funcName + " Settings",
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
			declarative.HSplitter{
				MaxSize: declarative.Size{Height: 30},
				Children: []declarative.Widget{
					declarative.HSpacer{},
					declarative.PushButton{
						AssignTo: &dialog.AcceptButton,
						Text:     "OK",
						Font:     declarative.Font{Family: "Roboto", PointSize: 9},
						OnClicked: func() {
							if err := dialog.DataBinder.Submit(); err != nil {
								fmt.Println(err)
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

// FillDialog Getting the right databinder & the right template needed
func FillDialog(funcName string) (interface{}, []declarative.Widget) {
	switch funcName {
	case "Sleep":
		return new(Sleep), SleepTemplate
	case "MilliSleep":
		return new(MilliSleep), MilliSleepTemplate
	case "Click":
		return new(MouseClick), MouseClickTemplate
	case "Scroll":
		return new(MouseScroll), MouseScrollTemplate
	case "Move":
		return new(MouseMove), MouseMoveTemplate
	case "Tap":
		return new(KeyboardTap), KeyboardTapTemplate
	case "Write":
		return new(ClipboardWrite), ClipboardWriteTemplate
	case "Read":
		return new(ClipboardRead), ClipboardReadTemplate
	case "Print":
		return new(LogPrint), LogPrintTemplate
	case "Create":
		return new(NotificationCreate), NotificationCreateTemplate
	case "GetBattery":
		return new(UserBattery), UserBatteryTemplate
	case "GetBatteryState":
		return new(BatteryState), BatteryParametersTemplate
	case "GetBatteryPercentage":
		return new(BatteryPercentage), BatteryParametersTemplate
	case "GetBatteryRemainingTime":
		return new(BatteryRemainingTime), BatteryParametersTemplate
	case "GetBatteryChargeRate":
		return new(BatteryChargeRate), BatteryParametersTemplate
	case "GetBatteryCurrentCapacity":
		return new(BatteryCurrentCapacity), BatteryParametersTemplate
	case "GetBatteryLastFullCapacity":
		return new(BatteryLastFullCapacity), BatteryParametersTemplate
	case "GetBatteryDesignCapacity":
		return new(BatteryDesignCapacity), BatteryParametersTemplate
	case "GetBatteryVoltage":
		return new(BatteryVoltage), BatteryParametersTemplate
	case "GetBatteryDesignVoltage":
		return new(BatteryDesignVoltage), BatteryParametersTemplate
	case "GetCurrentSysClock":
		return new(SysClock), SysClockTemplate
	case "GetCurrentSysTime":
		return new(SysTime), SysClockTemplate
	case "GetPixelColor":
		return new(PixelColor), PixelColorTemplate
	case "GetMouseColor":
		return new(MouseColor), MouseColorTemplate
	case "SaveCapture":
		return new(SaveCapture), SaveCaptureTemplate
	default:
		return nil, nil
	}
}
