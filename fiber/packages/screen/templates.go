package screen

import "github.com/lxn/walk/declarative"

// PixelColorTemplate Dialog's PixelColor Template
var PixelColorTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "X",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.NumberEdit{
				Value:    declarative.Bind("X"),
				Suffix:   " px",
				Decimals: 0,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("XIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Y",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.NumberEdit{
				Value:    declarative.Bind("Y"),
				Suffix:   " px",
				Decimals: 0,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("YIsVar"),
			},
		},
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
	declarative.GroupBox{
		Title:  "Path",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("Path"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("PathIsVar"),
			},
		},
	},
}
