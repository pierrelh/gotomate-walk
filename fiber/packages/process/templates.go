package process

import "github.com/lxn/walk/declarative"

// StartProcessTemplate Dialog's StartProcess Template
var StartProcessTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Path: ",
	},
	declarative.TextEdit{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Text:          declarative.Bind("Path"),
		CompactHeight: true,
	},
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Process Output: ",
	},
	declarative.TextEdit{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}
