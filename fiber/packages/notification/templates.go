package notification

import "github.com/lxn/walk/declarative"

// CreateTemplate Dialog's NotificationCreate Template
var CreateTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Title:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Title"),
		CompactHeight: true,
	},
	declarative.Label{
		Text: "Message:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Message"),
		CompactHeight: true,
	},
}
