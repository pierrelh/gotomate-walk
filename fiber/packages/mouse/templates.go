package mouse

import "github.com/lxn/walk/declarative"

// ClickTemplate Dialog's MouseClic Template
var ClickTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Click:",
	},
	declarative.ComboBox{
		Value:         declarative.Bind("MouseButtonName", declarative.SelRequired{}),
		BindingMember: "ID",
		DisplayMember: "Name",
		Model:         Buttons(),
	},
}

// ScrollTemplate Dialog's MouseScroll Template
var ScrollTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Horizontal:",
	},
	declarative.NumberEdit{
		Value:    declarative.Bind("X"),
		Decimals: 2,
	},
	declarative.Label{
		Text: "Vertical:",
	},
	declarative.NumberEdit{
		Value:    declarative.Bind("Y"),
		Decimals: 2,
	},
}

// MoveTemplate Dialog's MouseMove Template
var MoveTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Horizontal:",
	},
	declarative.NumberEdit{
		Value:    declarative.Bind("X"),
		Suffix:   " px",
		Decimals: 0,
	},
	declarative.Label{
		Text: "Vertical:",
	},
	declarative.NumberEdit{
		Value:    declarative.Bind("Y"),
		Suffix:   " px",
		Decimals: 0,
	},
}
