package screen

import "github.com/lxn/walk/declarative"

// PixelColorTemplate Dialog's PixelColor Template
var PixelColorTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "X: ",
	},
	declarative.NumberEdit{
		Font:     declarative.Font{Family: "Roboto", PointSize: 9},
		Value:    declarative.Bind("X"),
		Suffix:   " px",
		Decimals: 0,
	},
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Y: ",
	},
	declarative.NumberEdit{
		Font:     declarative.Font{Family: "Roboto", PointSize: 9},
		Value:    declarative.Bind("Y"),
		Suffix:   " px",
		Decimals: 0,
	},
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

// MouseColorTemplate Dialog's MouseColor Template
var MouseColorTemplate = []declarative.Widget{
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

// SaveCaptureTemplate Dialog's SaveCapture Template
var SaveCaptureTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Path: ",
	},
	declarative.TextEdit{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Text:          declarative.Bind("Path"),
		CompactHeight: true,
	},
}
