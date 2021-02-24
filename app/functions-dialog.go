package app

import (
	"fmt"
	"gotomate/packages"

	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
)

// CreateNewDialog Create the dialog for the function & add the needed template
func CreateNewDialog(funcName string) (interface{}, declarative.Dialog) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton
	source, children := FillDialog(funcName)

	return source,
		declarative.Dialog{
			Icon:          "/icon.ico",
			Title:         funcName + " Settings",
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
							Font:     declarative.Font{Family: "Roboto", PointSize: 9},
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
							Font:      declarative.Font{Family: "Roboto", PointSize: 9},
							OnClicked: func() { dlg.Cancel() },
						},
					},
				},
			},
		}
}

//FillDialog Getting the right databinder & the right template needed
func FillDialog(funcName string) (interface{}, []declarative.Widget) {
	switch funcName {
	case "Sleep":
		return new(packages.Sleep), packages.SleepTemplate
	case "MilliSleep":
		return new(packages.MilliSleep), packages.MilliSleepTemplate
	case "Click":
		return new(packages.MouseClick), packages.MouseClickTemplate
	case "Scroll":
		return new(packages.MouseScroll), packages.MouseScrollTemplate
	case "Move":
		return new(packages.MouseMove), packages.MouseMoveTemplate
	case "Tap":
		return new(packages.KeyboardTap), packages.KeyboardTapTemplate
	case "Write":
		return new(packages.ClipboardWrite), packages.ClipboardWriteTemplate
	case "Read":
		return new(packages.ClipboardRead), packages.ClipboardReadTemplate
	case "Print":
		return new(packages.LogPrint), packages.LogPrintTemplate
	case "Create":
		return new(packages.NotificationCreate), packages.NotificationCreateTemplate
	case "GetBattery":
		return new(packages.UserBattery), packages.UserBatteryTemplate
	case "GetBatteryState":
		return new(packages.BatteryState), packages.BatteryParametersTemplate
	case "GetBatteryPercentage":
		return new(packages.BatteryPercentage), packages.BatteryParametersTemplate
	case "GetBatteryRemainingTime":
		return new(packages.BatteryRemainingTime), packages.BatteryParametersTemplate
	case "GetBatteryChargeRate":
		return new(packages.BatteryChargeRate), packages.BatteryParametersTemplate
	case "GetBatteryCurrentCapacity":
		return new(packages.BatteryCurrentCapacity), packages.BatteryParametersTemplate
	case "GetBatteryLastFullCapacity":
		return new(packages.BatteryLastFullCapacity), packages.BatteryParametersTemplate
	case "GetBatteryDesignCapacity":
		return new(packages.BatteryDesignCapacity), packages.BatteryParametersTemplate
	case "GetBatteryVoltage":
		return new(packages.BatteryVoltage), packages.BatteryParametersTemplate
	case "GetBatteryDesignVoltage":
		return new(packages.BatteryDesignVoltage), packages.BatteryParametersTemplate
	case "GetCurrentSysClock":
		return new(packages.SysClock), packages.SysClockTemplate
	case "GetCurrentSysTime":
		return new(packages.SysTime), packages.SysClockTemplate
	default:
		return nil, nil
	}
}
