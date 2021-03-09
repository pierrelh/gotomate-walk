package systime

import "github.com/lxn/walk/declarative"

//SysClockTemplate Dialog's SysClock Template
var SysClockTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Output:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}
