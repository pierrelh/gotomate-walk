package battery

import "fmt"

// Decode Return the right databinder to decode a saved battery instruction
func Decode(function string) interface{} {
	switch function {
	case "GetBattery":
		return new(BatParameterDatabinder)
	case "GetBatteryState":
		return new(BatParameterDatabinder)
	case "GetBatteryPercentage":
		return new(BatParameterDatabinder)
	case "GetBatteryRemainingTime":
		return new(BatParameterDatabinder)
	case "GetBatteryChargeRate":
		return new(BatParameterDatabinder)
	case "GetBatteryCurrentCapacity":
		return new(BatParameterDatabinder)
	case "GetBatteryLastFullCapacity":
		return new(BatParameterDatabinder)
	case "GetBatteryDesignCapacity":
		return new(BatParameterDatabinder)
	case "GetBatteryVoltage":
		return new(BatParameterDatabinder)
	case "GetBatteryDesignVoltage":
		return new(BatParameterDatabinder)
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction decode")
	return nil
}
