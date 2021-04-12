package log

import "github.com/lxn/walk/declarative"

// PrintTemplate Dialog's LogPrint Template
var PrintTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Log",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("Log"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("LogIsVar"),
			},
		},
	},
}
