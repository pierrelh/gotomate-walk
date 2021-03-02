package systime

import "github.com/lxn/walk/declarative"

//SysClockTemplate Dialog's SysClock Template
var SysClockTemplate = []declarative.Widget{
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
