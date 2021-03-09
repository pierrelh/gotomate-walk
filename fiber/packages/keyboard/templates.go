package keyboard

import "github.com/lxn/walk/declarative"

// TapTemplate Dialog's KeyboardTap Template
var TapTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Input:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Input"),
		CompactHeight: true,
		MaxLength:     1,
	},
	declarative.Label{
		Text: "Special Input 1:",
	},
	declarative.ComboBox{
		Value:         declarative.Bind("Special1", declarative.SelRequired{}),
		BindingMember: "Name",
		DisplayMember: "Name",
		Model:         Inputs(),
	},
	declarative.Label{
		Text: "Special Input 2:",
	},
	declarative.ComboBox{
		Value:         declarative.Bind("Special2", declarative.SelRequired{}),
		BindingMember: "Name",
		DisplayMember: "Name",
		Model:         Inputs(),
	},
}
