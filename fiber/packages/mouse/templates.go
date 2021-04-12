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
	declarative.GroupBox{
		Title:  "Horizontal",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.NumberEdit{
				Value:    declarative.Bind("X"),
				Decimals: 2,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("XIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Vertical",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.NumberEdit{
				Value:    declarative.Bind("Y"),
				Decimals: 2,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("YIsVar"),
			},
		},
	},
}

// MoveTemplate Dialog's MouseMove Template
var MoveTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Horizontal",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.NumberEdit{
				Value:    declarative.Bind("X"),
				Suffix:   " px",
				Decimals: 0,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("XIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Vertical",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.NumberEdit{
				Value:    declarative.Bind("Y"),
				Suffix:   " px",
				Decimals: 0,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("YIsVar"),
			},
		},
	},
}
