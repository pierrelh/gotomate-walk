package clipboard

import "github.com/lxn/walk/declarative"

// WriteTemplate Dialog's ClipboardWrite Template
var WriteTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Clipboard",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text: declarative.Bind("Content"),
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("ContentIsVar"),
			},
		},
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
