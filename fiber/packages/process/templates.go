package process

import "github.com/lxn/walk/declarative"

// StartProcessTemplate Dialog's StartProcess Template
var StartProcessTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Path:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Path"),
		CompactHeight: true,
	},
	declarative.Label{
		Text: "Process Output:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}
