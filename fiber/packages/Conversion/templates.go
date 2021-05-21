package conversion

import "github.com/lxn/walk/declarative"

// IntToStringTemplate Dialog's IntToString Template
var IntToStringTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Input to convert",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("InputVarName"),
				Visible:       declarative.Bind("IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("Input"),
				Visible:  declarative.Bind("!IsAVar.Checked"),
				Decimals: 0,
			},
			declarative.CheckBox{
				Name:    "IsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("InputIsVar"),
			},
		},
	},
	declarative.Label{
		Text: "Output var:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}

// StringToBoolTemplate Dialog's StringToBool Template
var StringToBoolTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Input to convert",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("InputVarName"),
				Visible:       declarative.Bind("IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:          declarative.Bind("Input"),
				Visible:       declarative.Bind("!IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Name:    "IsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("InputIsVar"),
			},
		},
	},
	declarative.Label{
		Text: "Output var:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}

// StringToIntTemplate Dialog's StringToInt Template
var StringToIntTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Input to convert",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("InputVarName"),
				Visible:       declarative.Bind("IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:          declarative.Bind("Input"),
				Visible:       declarative.Bind("!IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Name:    "IsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("InputIsVar"),
			},
		},
	},
	declarative.Label{
		Text: "Output var:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}
