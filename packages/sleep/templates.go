package sleep

import "github.com/lxn/walk/declarative"

// SleepTemplate Dialog's Sleep Template
var SleepTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Sleep For:",
	},
	declarative.NumberEdit{
		Font:     declarative.Font{Family: "Roboto", PointSize: 9},
		Value:    declarative.Bind("Duration"),
		Suffix:   "s",
		Decimals: 0,
	},
}

// MilliSleepTemplate Dialog's MilliSleep Template
var MilliSleepTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Sleep For:",
	},
	declarative.NumberEdit{
		Font:     declarative.Font{Family: "Roboto", PointSize: 9},
		Value:    declarative.Bind("Duration"),
		Suffix:   "ms",
		Decimals: 0,
	},
}
