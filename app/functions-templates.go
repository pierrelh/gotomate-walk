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

//MouseClickTemplate Dialog's MouseClicTemplate Template
var MouseClickTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Click:",
	},
	declarative.ComboBox{
		Value:         declarative.Bind("MouseButtonName", declarative.SelRequired{}),
		BindingMember: "ID",
		DisplayMember: "Name",
		Model:         MouseButtons(),
	},
}
