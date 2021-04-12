package sleep

import "github.com/lxn/walk/declarative"

// SleepTemplate Dialog's Sleep Template
var SleepTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Sleep For",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.NumberEdit{
				Value:    declarative.Bind("Duration"),
				Suffix:   " s",
				Decimals: 0,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("DurationIsVar"),
			},
		},
	},
}

// MilliSleepTemplate Dialog's MilliSleep Template
var MilliSleepTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Sleep For",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.NumberEdit{
				Value:    declarative.Bind("Duration"),
				Suffix:   " ms",
				Decimals: 0,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("DurationIsVar"),
			},
		},
	},
}
