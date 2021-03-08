package packages

import (
	"fmt"
	"gotomate/fiber/packages/battery"
	"gotomate/fiber/packages/clipboard"
	"gotomate/fiber/packages/input"
	"gotomate/fiber/packages/keyboard"
	"gotomate/fiber/packages/log"
	"gotomate/fiber/packages/mouse"
	"gotomate/fiber/packages/notification"
	"gotomate/fiber/packages/process"
	"gotomate/fiber/packages/screen"
	"gotomate/fiber/packages/sleep"
	"gotomate/fiber/packages/systime"

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
		return new(sleep.Databinder), sleep.SleepTemplate
	case "MilliSleep":
		return new(sleep.Databinder), sleep.MilliSleepTemplate
	case "Click":
		return new(mouse.ClickDatabinder), mouse.ClickTemplate
	case "Scroll":
		return new(mouse.ScrollDatabinder), mouse.ScrollTemplate
	case "Move":
		return new(mouse.MoveDatabinder), mouse.MoveTemplate
	case "Tap":
		return new(keyboard.TapDatabinder), keyboard.TapTemplate
	case "Write":
		return new(clipboard.WriteDatabinder), clipboard.WriteTemplate
	case "Read":
		return new(clipboard.ReadDatabinder), clipboard.ReadTemplate
	case "Print":
		return new(log.PrintDatabinder), log.PrintTemplate
	case "Create":
		return new(notification.CreateDatabinder), notification.CreateTemplate
	case "GetBattery":
		return new(battery.BatParameterDatabinder), battery.UserBatteryTemplate
	case "GetBatteryState":
		return new(battery.BatParameterDatabinder), battery.ParametersTemplate
	case "GetBatteryPercentage":
		return new(battery.BatParameterDatabinder), battery.ParametersTemplate
	case "GetBatteryRemainingTime":
		return new(battery.BatParameterDatabinder), battery.ParametersTemplate
	case "GetBatteryChargeRate":
		return new(battery.BatParameterDatabinder), battery.ParametersTemplate
	case "GetBatteryCurrentCapacity":
		return new(battery.BatParameterDatabinder), battery.ParametersTemplate
	case "GetBatteryLastFullCapacity":
		return new(battery.BatParameterDatabinder), battery.ParametersTemplate
	case "GetBatteryDesignCapacity":
		return new(battery.BatParameterDatabinder), battery.ParametersTemplate
	case "GetBatteryVoltage":
		return new(battery.BatParameterDatabinder), battery.ParametersTemplate
	case "GetBatteryDesignVoltage":
		return new(battery.BatParameterDatabinder), battery.ParametersTemplate
	case "GetCurrentSysClock":
		return new(systime.ClockDatabinder), systime.SysClockTemplate
	case "GetCurrentSysTime":
		return new(systime.TimeDatabinder), systime.SysClockTemplate
	case "GetPixelColor":
		return new(screen.PixelColorDatabinder), screen.PixelColorTemplate
	case "GetMouseColor":
		return new(screen.MouseColorDatabinder), screen.MouseColorTemplate
	case "SaveCapture":
		return new(screen.SaveCaptureDatabinder), screen.SaveCaptureTemplate
	case "StartProcess":
		return new(process.StartProcessDatabinder), process.StartProcessTemplate
	case "String":
		return new(input.StringDatabinder), input.InputTemplate
	case "Int":
		return new(input.IntDatabinder), input.InputTemplate
	case "Bool":
		return new(input.BoolDatabinder), input.InputTemplate
	default:
		return nil, nil
	}
}
