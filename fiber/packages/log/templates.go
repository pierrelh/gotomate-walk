package log

import "github.com/lxn/walk/declarative"

// PrintTemplate Dialog's LogPrint Template
var PrintTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Log:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Log"),
		CompactHeight: true,
	},
}
