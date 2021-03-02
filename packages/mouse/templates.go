package mouse

import "github.com/lxn/walk/declarative"

// ClickTemplate Dialog's MouseClic Template
var ClickTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Click:",
	},
	declarative.ComboBox{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Value:         declarative.Bind("MouseButtonName", declarative.SelRequired{}),
		BindingMember: "ID",
		DisplayMember: "Name",
		Model:         Buttons(),
	},
}

// ScrollTemplate Dialog's MouseScroll Template
var ScrollTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Horizontal: ",
	},
	declarative.NumberEdit{
		Font:     declarative.Font{Family: "Roboto", PointSize: 9},
		Value:    declarative.Bind("X"),
		Decimals: 2,
	},
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Vertical: ",
	},
	declarative.NumberEdit{
		Font:     declarative.Font{Family: "Roboto", PointSize: 9},
		Value:    declarative.Bind("Y"),
		Decimals: 2,
	},
}

// MoveTemplate Dialog's MouseMove Template
var MoveTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Horizontal: ",
	},
	declarative.NumberEdit{
		Font:     declarative.Font{Family: "Roboto", PointSize: 9},
		Value:    declarative.Bind("X"),
		Suffix:   "px",
		Decimals: 0,
	},
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Vertical: ",
	},
	declarative.NumberEdit{
		Font:     declarative.Font{Family: "Roboto", PointSize: 9},
		Value:    declarative.Bind("Y"),
		Suffix:   "px",
		Decimals: 0,
	},
}
