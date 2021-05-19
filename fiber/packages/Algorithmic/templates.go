package algorithmic

import "github.com/lxn/walk/declarative"

// IfTemplate Dialog's If Template
var IfTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Value 1:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("VarOneName"),
				Visible:       declarative.Bind("OneIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:          declarative.Bind("ValueOne"),
				Visible:       declarative.Bind("!OneIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Name:    "OneIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("OneIsVar"),
			},
		},
	},
	declarative.Label{
		Text: "Comparator:",
	},
	declarative.ComboBox{
		Value:         declarative.Bind("Comparator", declarative.SelRequired{}),
		BindingMember: "Name",
		DisplayMember: "Name",
		Model:         Comparators(),
	},
	declarative.GroupBox{
		Title:  "Value 2:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("VarTwoName"),
				Visible:       declarative.Bind("TwoIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:          declarative.Bind("ValueTwo"),
				Visible:       declarative.Bind("!TwoIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Name:    "TwoIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("TwoIsVar"),
			},
		},
	},
	declarative.Label{
		Text: "Else instruction ID:",
	},
	declarative.NumberEdit{
		Value:    declarative.Bind("FalseInstruction"),
		Decimals: 0,
	},
}
