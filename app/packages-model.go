package app

// Packages Export all the available automate's packages
var Packages = []string{
	"Battery",
	"Clipboard",
	"Keyboard",
	"Log",
	"Mouse",
	"Notification",
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
	"Sleep": {
		"MilliSleep",
		"Sleep",
	},
	"Systime": {
		"GetCurrentSysTime",
		"GetCurrentSysClock",
	},
}
