package automate

// Packages Export all the available automate's packages
var Packages = []string{
	"Battery",
	"Clipboard",
	"Keyboard",
	"Log",
	"Mouse",
	"Notification",
	"Sleep",
}

// SubPackages Export all the available functions
var SubPackages = map[string][]string{
	"Battery": {
		"GetBatteries",
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
		"ReadClipboard",
		"WriteClipboard",
	},
	"Keyboard": {
		"KeyTap",
	},
	"Log": {
		"Print",
	},
	"Mouse": {
		"MouseMove",
	},
	"Notification": {
		"CreateNotification",
		"PushNotification",
	},
	"Sleep": {
		"MilliSleep",
		"Sleep",
	},
}
