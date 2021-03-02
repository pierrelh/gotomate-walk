package keyboard

import "github.com/lxn/walk/declarative"

// TapTemplate Dialog's KeyboardTap Template
var TapTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Input: ",
	},
	declarative.TextEdit{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Text:          declarative.Bind("Input"),
		CompactHeight: true,
		MaxLength:     1,
	},
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Special Input 1:",
	},
	declarative.ComboBox{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Value:         declarative.Bind("Special1", declarative.SelRequired{}),
		BindingMember: "Name",
		DisplayMember: "Name",
		Model:         Inputs(),
	},
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Special Input 2:",
	},
	declarative.ComboBox{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Value:         declarative.Bind("Special2", declarative.SelRequired{}),
		BindingMember: "Name",
		DisplayMember: "Name",
		Model:         Inputs(),
	},
}
