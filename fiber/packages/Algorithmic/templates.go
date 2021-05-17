package algorithmic

import "github.com/lxn/walk/declarative"

// DefineTemplate Dialog's Define Template
var DefineTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Name:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Name"),
		CompactHeight: true,
	},
	declarative.GroupBox{
		Title:  "Value:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("Value"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("IsVar"),
			},
		},
	},
}

// IfTemplate Dialog's If Template
var IfTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Value 1:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ValueOne"),
				CompactHeight: true,
			},
			declarative.CheckBox{
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
				Text:          declarative.Bind("ValueTwo"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("TwoIsVar"),
			},
		},
	},
	declarative.Label{
		Text: "If false instruction ID:",
	},
	declarative.NumberEdit{
		Value:    declarative.Bind("FalseInstruction"),
		Decimals: 0,
	},
}
