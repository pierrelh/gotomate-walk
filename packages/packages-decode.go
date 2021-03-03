package packages

import (
	"fmt"
	"gotomate/app/fiber"
	"gotomate/packages/battery"
	"gotomate/packages/clipboard"
	"gotomate/packages/input"
	"gotomate/packages/keyboard"
	"gotomate/packages/log"
	"gotomate/packages/mouse"
	"gotomate/packages/notification"
	"gotomate/packages/process"
	"gotomate/packages/screen"
	"gotomate/packages/sleep"
	"gotomate/packages/systime"
)

// PackageDecode return the struct to use to decode an unmarshalled data
func PackageDecode(instruction *fiber.LoadingInstruction) interface{} {
	var structure interface{}
	switch instruction.FuncName {
	case "Sleep":
		structure = new(sleep.SleepDatabinder)
	case "MilliSleep":
		structure = new(sleep.MilliSleepDatabinder)
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
		structure = new(battery.UserBatteryDatabinder)
	case "GetBatteryState":
		structure = new(battery.StateDatabinder)
	case "GetBatteryPercentage":
		structure = new(battery.PercentageDatabinder)
	case "GetBatteryRemainingTime":
		structure = new(battery.RemainingTimeDatabinder)
	case "GetBatteryChargeRate":
		structure = new(battery.ChargeRateDatabinder)
	case "GetBatteryCurrentCapacity":
		structure = new(battery.CurrentCapacityDatabinder)
	case "GetBatteryLastFullCapacity":
		structure = new(battery.LastFullCapacityDatabinder)
	case "GetBatteryDesignCapacity":
		structure = new(battery.DesignCapacityDatabinder)
	case "GetBatteryVoltage":
		structure = new(battery.VoltageDatabinder)
	case "GetBatteryDesignVoltage":
		structure = new(battery.DesignVoltageDatabinder)
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
