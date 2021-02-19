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
		Value:    declarative.Bind("Duration"),
		Suffix:   "s",
		Decimals: 0,
	},
}

//MilliSleepTemplate Dialog's MilliSleep Template
var MilliSleepTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Sleep For:",
	},
	declarative.NumberEdit{
		Value:    declarative.Bind("Duration"),
		Suffix:   "ms",
		Decimals: 0,
	},
}

//MouseClickTemplate Dialog's MouseClic Template
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

//MouseScrollTemplate Dialog's MouseScroll Template
var MouseScrollTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Horizontal: ",
	},
	declarative.NumberEdit{
		Value:    declarative.Bind("X"),
		Decimals: 2,
	},
	declarative.Label{
		Text: "Vertical: ",
	},
	declarative.NumberEdit{
		Value:    declarative.Bind("Y"),
		Decimals: 2,
	},
}

//MouseMoveTemplate Dialog's MouseMove Template
var MouseMoveTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Horizontal: ",
	},
	declarative.NumberEdit{
		Value:    declarative.Bind("X"),
		Suffix:   "px",
		Decimals: 0,
	},
	declarative.Label{
		Text: "Vertical: ",
	},
	declarative.NumberEdit{
		Value:    declarative.Bind("Y"),
		Suffix:   "px",
		Decimals: 0,
	},
}
