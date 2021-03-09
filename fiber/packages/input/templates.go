package input

import "github.com/lxn/walk/declarative"

// InputTemplate Dialog's Input Template
var InputTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Message:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Message"),
		CompactHeight: true,
	},
	declarative.Label{
		Text: "Output var:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}
