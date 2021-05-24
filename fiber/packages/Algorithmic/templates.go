package algorithmic

import "github.com/lxn/walk/declarative"

// ForTemplate Dialog's For Template
var ForTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Variable to increment",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("VarOneName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
	declarative.ComboBox{
		Value:         declarative.Bind("Comparator", declarative.SelRequired{}),
		BindingMember: "Name",
		DisplayMember: "Name",
		Model:         Comparators(),
	},
	declarative.GroupBox{
		Title:  "Value to compare",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("VarTwoName"),
				Visible:       declarative.Bind("TwoIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:   declarative.Bind("ValueTwo"),
				Visible: declarative.Bind("!TwoIsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "TwoIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("TwoIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Increment",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("IncrementVarName"),
				Visible:       declarative.Bind("IncrementIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:   declarative.Bind("Increment"),
				Visible: declarative.Bind("!IncrementIsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "IncrementIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("IncrementIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Else instruction ID",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("FalseInstructionVarName"),
				Visible:       declarative.Bind("FalseInstructionIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("FalseInstruction"),
				Visible:  declarative.Bind("!FalseInstructionIsAVar.Checked"),
				Decimals: 0,
			},
			declarative.CheckBox{
				Name:    "FalseInstructionIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("FalseInstructionIsVar"),
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
	declarative.GroupBox{
		Title:  "Else instruction ID",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("FalseInstructionVarName"),
				Visible:       declarative.Bind("FalseInstructionIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("FalseInstruction"),
				Visible:  declarative.Bind("!FalseInstructionIsAVar.Checked"),
				Decimals: 0,
			},
			declarative.CheckBox{
				Name:    "FalseInstructionIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("FalseInstructionIsVar"),
			},
		},
	},
}
