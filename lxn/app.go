package lxn

import (
	"piproto/automate"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// CreateApp Initiate the app
func CreateApp() {
	var inTE, outTE *walk.TextEdit
	/*
	   Use:
	   Databinding
	   listbox
	   log view
	   multipage
	   notifyicon
	*/

	MainWindow{
		Icon:    "img/icon.ico",
		Title:   "Gotomate",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text:      "RUN",
				OnClicked: automate.LaunchAutomate,
			},
		},
	}.Run()
}
