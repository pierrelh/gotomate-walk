package packages

import (
	"fmt"
	"gotomate/fiber"
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
)

// PackageDecode return the struct to use to decode an unmarshalled data
func PackageDecode(instruction *fiber.LoadingInstruction) interface{} {
	var structure interface{}
	switch instruction.FuncName {
	case "Sleep":
		structure = new(sleep.Databinder)
	case "MilliSleep":
		structure = new(sleep.Databinder)
	case "Click":
		structure = new(mouse.ClickDatabinder)
	case "Scroll":
		structure = new(mouse.ScrollDatabinder)
	case "Move":
		structure = new(mouse.MoveDatabinder)
	case "Tap":
		structure = new(keyboard.TapDatabinder)
	case "Write":
		structure = new(clipboard.WriteDatabinder)
	case "Read":
		structure = new(clipboard.ReadDatabinder)
	case "Print":
		structure = new(log.PrintDatabinder)
	case "Create":
		structure = new(notification.CreateDatabinder)
	case "GetBattery":
		structure = new(battery.BatParameterDatabinder)
	case "GetBatteryState":
		structure = new(battery.BatParameterDatabinder)
	case "GetBatteryPercentage":
		structure = new(battery.BatParameterDatabinder)
	case "GetBatteryRemainingTime":
		structure = new(battery.BatParameterDatabinder)
	case "GetBatteryChargeRate":
		structure = new(battery.BatParameterDatabinder)
	case "GetBatteryCurrentCapacity":
		structure = new(battery.BatParameterDatabinder)
	case "GetBatteryLastFullCapacity":
		structure = new(battery.BatParameterDatabinder)
	case "GetBatteryDesignCapacity":
		structure = new(battery.BatParameterDatabinder)
	case "GetBatteryVoltage":
		structure = new(battery.BatParameterDatabinder)
	case "GetBatteryDesignVoltage":
		structure = new(battery.BatParameterDatabinder)
	case "GetCurrentSysClock":
		structure = new(systime.ClockDatabinder)
	case "GetCurrentSysTime":
		structure = new(systime.TimeDatabinder)
	case "GetPixelColor":
		structure = new(screen.PixelColorDatabinder)
	case "GetMouseColor":
		structure = new(screen.MouseColorDatabinder)
	case "SaveCapture":
		structure = new(screen.SaveCaptureDatabinder)
	case "StartProcess":
		structure = new(process.StartProcessDatabinder)
	case "String":
		structure = new(input.StringDatabinder)
	case "Int":
		structure = new(input.IntDatabinder)
	case "Bool":
		structure = new(input.BoolDatabinder)
	default:
		fmt.Println("ERROR: Unable to find the function")
		structure = nil
	}
	return structure
}
