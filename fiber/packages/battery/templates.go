package battery

import "github.com/lxn/walk/declarative"

// UserBatteryTemplate Dialog's UserBattery Template
var UserBatteryTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Output var:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}

// ParametersTemplate Dialog's BatteryParameters Template
var ParametersTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Battery Name:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("BatteryName"),
		CompactHeight: true,
	},
	declarative.Label{
		Text: "Output var:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}
