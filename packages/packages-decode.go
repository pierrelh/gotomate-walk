package packages

import (
	"fmt"
	"gotomate/app/fiber"
)

//PackageDecode return the struct to use to decode an unmarshalled data
func PackageDecode(instruction *fiber.LoadingInstruction) interface{} {
	var structure interface{}
	switch instruction.FuncName {
	case "Log":
		structure = new(LogPrint)
	case "Sleep":
		structure = new(Sleep)
	case "MilliSleep":
		structure = new(MilliSleep)
	case "Click":
		structure = new(MouseClick)
	case "Scroll":
		structure = new(MouseScroll)
	case "Move":
		structure = new(MouseMove)
	case "Tap":
		structure = new(KeyboardTap)
	case "Write":
		structure = new(ClipboardWrite)
	case "Read":
		structure = new(ClipboardRead)
	case "Print":
		structure = new(LogPrint)
	case "Create":
		structure = new(NotificationCreate)
	case "GetBattery":
		structure = new(UserBattery)
	case "GetBatteryState":
		structure = new(BatteryState)
	case "GetBatteryPercentage":
		structure = new(BatteryPercentage)
	case "GetBatteryRemainingTime":
		structure = new(BatteryRemainingTime)
	case "GetBatteryChargeRate":
		structure = new(BatteryChargeRate)
	case "GetBatteryCurrentCapacity":
		structure = new(BatteryCurrentCapacity)
	case "GetBatteryLastFullCapacity":
		structure = new(BatteryLastFullCapacity)
	case "GetBatteryDesignCapacity":
		structure = new(BatteryDesignCapacity)
	case "GetBatteryVoltage":
		structure = new(BatteryVoltage)
	case "GetBatteryDesignVoltage":
		structure = new(BatteryDesignVoltage)
	case "GetCurrentSysClock":
		structure = new(SysClock)
	case "GetCurrentSysTime":
		structure = new(SysTime)
	default:
		fmt.Println("ERROR: Unable to find the function")
		structure = nil
	}
	return structure
}
