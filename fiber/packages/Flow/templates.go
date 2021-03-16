package flow

import "github.com/lxn/walk/declarative"

// StartTemplate Dialog's ClipboardWrite Template
var StartTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Flow Name:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("FlowName"),
		CompactHeight: true,
	},
}
