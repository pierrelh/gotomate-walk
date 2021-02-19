package app

import (
	. "github.com/lxn/walk/declarative"
)

//SleepTemplate Dialog's Sleep Template
var SleepTemplate = []Widget{
	Label{
		Text: "Sleep For:",
	},
	NumberEdit{
		Value:    Bind("Duration", Range{Min: 1, Max: 1000}),
		Suffix:   " s",
		Decimals: 0,
	},
}

//MilliSleepTemplate Dialog's MilliSleep Template
var MilliSleepTemplate = []Widget{
	Label{
		Text: "Sleep For:",
	},
	NumberEdit{
		Value:    Bind("Duration", Range{Min: 1, Max: 1000}),
		Suffix:   " ms",
		Decimals: 0,
	},
}
