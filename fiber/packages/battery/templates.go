package battery

import "github.com/lxn/walk/declarative"

// UserBatteryTemplate Dialog's UserBattery Template
var UserBatteryTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Output var: ",
	},
	declarative.TextEdit{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}

// ParametersTemplate Dialog's BatteryParameters Template
var ParametersTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Battery Name: ",
	},
	declarative.TextEdit{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Text:          declarative.Bind("BatteryName"),
		CompactHeight: true,
	},
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Output var: ",
	},
	declarative.TextEdit{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}
