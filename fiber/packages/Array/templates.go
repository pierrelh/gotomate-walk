package array

import "github.com/lxn/walk/declarative"

// PopAtTemplate Dialog's PopAt Template
var PopAtTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Array:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ArrayVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
	declarative.GroupBox{
		Title:  "Index to pop:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("IndexVarName"),
				Visible:       declarative.Bind("IsIndexAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("Index"),
				Visible:  declarative.Bind("!IsIndexAVar.Checked"),
				Decimals: 0,
			},
			declarative.CheckBox{
				Name:    "IsIndexAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("IndexIsVar"),
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

// PopLastTemplate Dialog's PopLast Template
var PopLastTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Array:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ArrayVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
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

// PushAtTemplate Dialog's PushAt Template
var PushAtTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Array:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ArrayVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
	declarative.GroupBox{
		Title:  "Index to push at:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("IndexVarName"),
				Visible:       declarative.Bind("IsIndexAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("Index"),
				Visible:  declarative.Bind("!IsIndexAVar.Checked"),
				Decimals: 0,
			},
			declarative.CheckBox{
				Name:    "IsIndexAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("IndexIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Value to push:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ValueVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
}

// PushLastTemplate Dialog's PushLast Template
var PushLastTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Array:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ArrayVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
	declarative.GroupBox{
		Title:  "Value to push:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ValueVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
}

// RemoveAtTemplate Dialog's RemoveAt Template
var RemoveAtTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Array:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ArrayVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
	declarative.GroupBox{
		Title:  "Index to remove:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("IndexVarName"),
				Visible:       declarative.Bind("IsIndexAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("Index"),
				Visible:  declarative.Bind("!IsIndexAVar.Checked"),
				Decimals: 0,
			},
			declarative.CheckBox{
				Name:    "IsIndexAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("IndexIsVar"),
			},
		},
	},
}

// RemoveLastTemplate Dialog's RemoveLast Template
var RemoveLastTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Array:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ArrayVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
}

// UpdateValueTemplate Dialog's UpdateValue Template
var UpdateValueTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Array:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ArrayVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
	declarative.GroupBox{
		Title:  "Index to update:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("IndexVarName"),
				Visible:       declarative.Bind("IsIndexAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("Index"),
				Visible:  declarative.Bind("!IsIndexAVar.Checked"),
				Decimals: 0,
			},
			declarative.CheckBox{
				Name:    "IsIndexAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("IndexIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "New value:",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ValueVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
}
