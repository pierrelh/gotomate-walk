package clipboard

import "github.com/lxn/walk/declarative"

// WriteTemplate Dialog's ClipboardWrite Template
var WriteTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Clipboard:",
	},
	declarative.TextEdit{
		Text: declarative.Bind("Content"),
	},
}

// ReadTemplate Dialog's ClipboardRead Template
var ReadTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Output var:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}
