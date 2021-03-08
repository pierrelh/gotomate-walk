package packages

// Packages Export all the available automate's packages
var Packages = []string{
	"Battery",
	"Clipboard",
	"Input",
	"Keyboard",
	"Log",
	"Mouse",
	"Notification",
	"Process",
	"Screen",
	"Sleep",
	"Systime",
}

// SubPackages Export all the available functions
var SubPackages = map[string][]string{
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
