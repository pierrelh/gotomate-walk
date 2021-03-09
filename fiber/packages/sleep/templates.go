package sleep

import "github.com/lxn/walk/declarative"

// SleepTemplate Dialog's Sleep Template
var SleepTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Sleep For:",
	},
	declarative.NumberEdit{
		Value:    declarative.Bind("Duration"),
		Suffix:   " s",
		Decimals: 0,
	},
}

// MilliSleepTemplate Dialog's MilliSleep Template
var MilliSleepTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Sleep For:",
	},
	declarative.NumberEdit{
		Value:    declarative.Bind("Duration"),
		Suffix:   " ms",
		Decimals: 0,
	},
}
