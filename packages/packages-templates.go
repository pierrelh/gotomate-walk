package packages

import "github.com/lxn/walk/declarative"

//SleepTemplate Dialog's Sleep Template
var SleepTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Sleep For:",
	},
	declarative.NumberEdit{
		Font:     declarative.Font{Family: "Roboto", PointSize: 9},
		Value:    declarative.Bind("Duration"),
		Suffix:   "s",
		Decimals: 0,
	},
}

//MilliSleepTemplate Dialog's MilliSleep Template
var MilliSleepTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Sleep For:",
	},
	declarative.NumberEdit{
		Font:     declarative.Font{Family: "Roboto", PointSize: 9},
		Value:    declarative.Bind("Duration"),
		Suffix:   "ms",
		Decimals: 0,
	},
}

//MouseClickTemplate Dialog's MouseClic Template
var MouseClickTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Click:",
	},
	declarative.ComboBox{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Value:         declarative.Bind("MouseButtonName", declarative.SelRequired{}),
		BindingMember: "ID",
		DisplayMember: "Name",
		Model:         MouseButtons(),
	},
}

//MouseScrollTemplate Dialog's MouseScroll Template
var MouseScrollTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Horizontal: ",
	},
	declarative.NumberEdit{
		Font:     declarative.Font{Family: "Roboto", PointSize: 9},
		Value:    declarative.Bind("X"),
		Decimals: 2,
	},
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Vertical: ",
	},
	declarative.NumberEdit{
		Font:     declarative.Font{Family: "Roboto", PointSize: 9},
		Value:    declarative.Bind("Y"),
		Decimals: 2,
	},
}

//MouseMoveTemplate Dialog's MouseMove Template
var MouseMoveTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Horizontal: ",
	},
	declarative.NumberEdit{
		Font:     declarative.Font{Family: "Roboto", PointSize: 9},
		Value:    declarative.Bind("X"),
		Suffix:   "px",
		Decimals: 0,
	},
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Vertical: ",
	},
	declarative.NumberEdit{
		Font:     declarative.Font{Family: "Roboto", PointSize: 9},
		Value:    declarative.Bind("Y"),
		Suffix:   "px",
		Decimals: 0,
	},
}

//KeyboardTapTemplate Dialog's KeyboardTap Template
var KeyboardTapTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Input: ",
	},
	declarative.TextEdit{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Text:          declarative.Bind("Input"),
		CompactHeight: true,
		MaxLength:     1,
	},
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Special Input 1:",
	},
	declarative.ComboBox{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Value:         declarative.Bind("Special1", declarative.SelRequired{}),
		BindingMember: "Name",
		DisplayMember: "Name",
		Model:         KeyboardInputs(),
	},
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Special Input 2:",
	},
	declarative.ComboBox{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Value:         declarative.Bind("Special2", declarative.SelRequired{}),
		BindingMember: "Name",
		DisplayMember: "Name",
		Model:         KeyboardInputs(),
	},
}

//ClipboardWriteTemplate Dialog's ClipboardWrite Template
var ClipboardWriteTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Clipboard: ",
	},
	declarative.TextEdit{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: declarative.Bind("Content"),
	},
}

//ClipboardReadTemplate Dialog's ClipboardRead Template
var ClipboardReadTemplate = []declarative.Widget{
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

//LogPrintTemplate Dialog's LogPrint Template
var LogPrintTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Log: ",
	},
	declarative.TextEdit{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Text:          declarative.Bind("Log"),
		CompactHeight: true,
	},
}

//NotificationCreateTemplate Dialog's NotificationCreate Template
var NotificationCreateTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Title: ",
	},
	declarative.TextEdit{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Text:          declarative.Bind("Title"),
		CompactHeight: true,
	},
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Message: ",
	},
	declarative.TextEdit{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Text:          declarative.Bind("Message"),
		CompactHeight: true,
	},
}

//UserBatteryTemplate Dialog's UserBattery Template
var UserBatteryTemplate = []declarative.Widget{
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

//BatteryParametersTemplate Dialog's BatteryParameters Template
var BatteryParametersTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Battery Name: ",
	},
	declarative.TextEdit{
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		Text:          declarative.Bind("BatteryName"),
		CompactHeight: true,
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

// PixelColorTemplate Dialog's PixelColor Template
var PixelColorTemplate = []declarative.Widget{
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "X: ",
	},
	declarative.NumberEdit{
		Font:     declarative.Font{Family: "Roboto", PointSize: 9},
		Value:    declarative.Bind("X"),
		Suffix:   "px",
		Decimals: 0,
	},
	declarative.Label{
		Font: declarative.Font{Family: "Roboto", PointSize: 9},
		Text: "Y: ",
	},
	declarative.NumberEdit{
		Font:     declarative.Font{Family: "Roboto", PointSize: 9},
		Value:    declarative.Bind("Y"),
		Suffix:   "px",
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
