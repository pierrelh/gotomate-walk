package notification

import "github.com/lxn/walk/declarative"

// CreateTemplate Dialog's NotificationCreate Template
var CreateTemplate = []declarative.Widget{
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
