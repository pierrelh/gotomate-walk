package arithmetic

import "github.com/lxn/walk/declarative"

// DivideTemplate Dialog's Divide Template
var DivideTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Value One:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ValueOne"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("ValueOneIsVar"),
			},
		},
	},
	declarative.Label{
		Text:          "÷",
		TextAlignment: declarative.AlignCenter,
		Font:          declarative.Font{Family: "Roboto", PointSize: 12, Bold: true},
	},
	declarative.GroupBox{
		Title:  "Value Two:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ValueTwo"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("ValueTwoIsVar"),
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

// MultiplyTemplate Dialog's Multiply Template
var MultiplyTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Value One:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ValueOne"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("ValueOneIsVar"),
			},
		},
	},
	declarative.Label{
		Text:          "X",
		TextAlignment: declarative.AlignCenter,
		Font:          declarative.Font{Family: "Roboto", PointSize: 12, Bold: true},
	},
	declarative.GroupBox{
		Title:  "Value Two:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ValueTwo"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("ValueTwoIsVar"),
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

// PowTemplate Dialog's Pow Template
var PowTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Value One:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ValueOne"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("ValueOneIsVar"),
			},
		},
	},
	declarative.Label{
		Text:          "Pow",
		TextAlignment: declarative.AlignCenter,
		Font:          declarative.Font{Family: "Roboto", PointSize: 12, Bold: true},
	},
	declarative.GroupBox{
		Title:  "Value Two:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ValueTwo"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("ValueTwoIsVar"),
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

// SqrtTemplate Dialog's Sqrt Template
var SqrtTemplate = []declarative.Widget{
	declarative.Label{
		Text:          "√",
		TextAlignment: declarative.AlignCenter,
		Font:          declarative.Font{Family: "Roboto", PointSize: 12, Bold: true},
	},
	declarative.GroupBox{
		Title:  "Value:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("Value"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("ValueIsVar"),
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

// SubstractTemplate Dialog's Substract Template
var SubstractTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Value One:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ValueOne"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("ValueOneIsVar"),
			},
		},
	},
	declarative.Label{
		Text:          "-",
		TextAlignment: declarative.AlignCenter,
		Font:          declarative.Font{Family: "Roboto", PointSize: 12, Bold: true},
	},
	declarative.GroupBox{
		Title:  "Value Two:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ValueTwo"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("ValueTwoIsVar"),
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

// SumTemplate Dialog's Sum Template
var SumTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Value One:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ValueOne"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("ValueOneIsVar"),
			},
		},
	},
	declarative.Label{
		Text:          "+",
		TextAlignment: declarative.AlignCenter,
		Font:          declarative.Font{Family: "Roboto", PointSize: 12, Bold: true},
	},
	declarative.GroupBox{
		Title:  "Value Two:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ValueTwo"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("ValueTwoIsVar"),
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
