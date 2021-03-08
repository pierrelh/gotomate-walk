package log

import "github.com/lxn/walk/declarative"

// PrintTemplate Dialog's LogPrint Template
var PrintTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Log:",
	},
	declarative.TextEdit{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Text:          declarative.Bind("Log"),
		CompactHeight: true,
	},
}
