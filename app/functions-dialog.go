package app

import (
	"fmt"

	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
)

// CreateNewDialog Create the dialog for the function & add the needed template
func CreateNewDialog(funcName string) declarative.Dialog {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton
	source, children := fillDialog(funcName)

	return declarative.Dialog{
		Icon:          "/icon.ico",
		Title:         "Settings",
		AssignTo:      &dlg,
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: declarative.DataBinder{
			AssignTo:       &db,
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
						AssignTo: &acceptPB,
						Text:     "OK",
						Font:     declarative.Font{Family: "Segoe UI", PointSize: 9},
						OnClicked: func() {
							if err := db.Submit(); err != nil {
								fmt.Println(err)
								return
							}

							dlg.Accept()
						},
					},
					declarative.PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						Font:      declarative.Font{Family: "Segoe UI", PointSize: 9},
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}
}

func fillDialog(funcName string) (interface{}, []declarative.Widget) {
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
	default:
		return nil, nil
	}
}
