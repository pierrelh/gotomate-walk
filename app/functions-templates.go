package app

import (
	declarative "github.com/lxn/walk/declarative"
)

//SleepTemplate Dialog's Sleep Template
var SleepTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Sleep For:",
	},
	declarative.NumberEdit{
		Value:    declarative.Bind("Duration", declarative.Range{Min: 1, Max: 1000}),
		Suffix:   " s",
		Decimals: 0,
	},
}

//MilliSleepTemplate Dialog's MilliSleep Template
var MilliSleepTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Sleep For:",
	},
	declarative.NumberEdit{
		Value:    declarative.Bind("Duration", declarative.Range{Min: 1, Max: 1000}),
		Suffix:   " ms",
		Decimals: 0,
	},
}
