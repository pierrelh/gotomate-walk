package app

import (
	. "github.com/lxn/walk/declarative"
)

var SleepTemplate = []Widget{
	Label{
		ColumnSpan: 2,
		Text:       "Sleep For:",
	},
	LineEdit{
		ColumnSpan: 2,
		Text:       Bind("SleepDurationField"),
	},
}
