package screen

import "github.com/lxn/walk/declarative"

// PixelColorTemplate Dialog's PixelColor Template
var PixelColorTemplate = []declarative.Widget{
	declarative.Label{
		Text: "X:",
	},
	declarative.NumberEdit{
		Value:    declarative.Bind("X"),
		Suffix:   " px",
		Decimals: 0,
	},
	declarative.Label{
		Text: "Y:",
	},
	declarative.NumberEdit{
		Value:    declarative.Bind("Y"),
		Suffix:   " px",
		Decimals: 0,
	},
	declarative.Label{
		Text: "Output:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}

// MouseColorTemplate Dialog's MouseColor Template
var MouseColorTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Output:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}

// ScreenSizeTemplate Dialog's ScreenSize Template
var ScreenSizeTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Height Output:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("HeightOutput"),
		CompactHeight: true,
	},
	declarative.Label{
		Text: "Width Output:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("WidthOutput"),
		CompactHeight: true,
	},
}

// SaveCaptureTemplate Dialog's SaveCapture Template
var SaveCaptureTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Path:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Path"),
		CompactHeight: true,
	},
}
