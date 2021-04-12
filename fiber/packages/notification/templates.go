package notification

import "github.com/lxn/walk/declarative"

// CreateTemplate Dialog's NotificationCreate Template
var CreateTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Title",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("Title"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("TitleIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Message",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("Message"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: declarative.Bind("MessageIsVar"),
			},
		},
	},
}
