package process

import "github.com/lxn/walk/declarative"

// StartProcessTemplate Dialog's StartProcess Template
var StartProcessTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Path",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text: declarative.Bind("Path"),
			},
			declarative.CheckBox{
				Text:      "Is a Var",
				Alignment: declarative.AlignHFarVCenter,
				Checked:   declarative.Bind("PathIsVar"),
			},
		},
	},
	declarative.Label{
		Text: "PID Output:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}

// KillProcessTemplate Dialog's KillProcess Template
var KillProcessTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "PID of process to kill",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("PID"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("PIDIsVar"),
			},
		},
	},
}
