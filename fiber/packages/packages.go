package packages

// Packages Export all the available packages & their functions
var Packages = map[string][]string{
	"Battery": {
		"GetBattery",
		"GetBatteryState",
		"GetBatteryPercentage",
		"GetBatteryRemainingTime",
		"GetBatteryChargeRate",
		"GetBatteryCurrentCapacity",
		"GetBatteryLastFullCapacity",
		"GetBatteryDesignCapacity",
		"GetBatteryVoltage",
		"GetBatteryDesignVoltage",
	},
	"Clipboard": {
		"Read",
		"Write",
	},
	"Input": {
		"Bool",
		"Int",
		"String",
	},
	"Keyboard": {
		"Tap",
	},
	"Log": {
		"Print",
	},
	"Mouse": {
		"Click",
		"Move",
		"Scroll",
	},
	"Notification": {
		"Create",
	},
	"Process": {
		"StartProcess",
	},
	"Screen": {
		"GetMouseColor",
		"GetPixelColor",
		// "GetScreenSize",
		"SaveCapture",
	},
	"Sleep": {
		"MilliSleep",
		"Sleep",
	},
	"Systime": {
		"GetCurrentSysTime",
		"GetCurrentSysClock",
	},
}
