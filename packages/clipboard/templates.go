package clipboard

import "github.com/lxn/walk/declarative"

// WriteTemplate Dialog's ClipboardWrite Template
var WriteTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Clipboard: ",
	},
	declarative.TextEdit{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: declarative.Bind("Content"),
	},
}

// ReadTemplate Dialog's ClipboardRead Template
var ReadTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Output var: ",
	},
	declarative.TextEdit{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Text:          declarative.Bind("Var"),
		CompactHeight: true,
	},
}
