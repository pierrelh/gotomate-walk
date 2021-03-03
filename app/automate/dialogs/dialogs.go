package dialogs

import (
	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
)

// Dlg Declare a new dialog
var Dlg *walk.Dialog

var acceptPB *walk.PushButton
var cancelPB *walk.PushButton

// FiberNotSavedDialog This dialog is showed when a not saved fiber is about to be deleted / replaced
var FiberNotSavedDialog = declarative.Dialog{
	Icon:          "/icon.ico",
	Title:         "Error",
	AssignTo:      &Dlg,
	DefaultButton: &cancelPB,
	MinSize: declarative.Size{
		Width:  200,
		Height: 150,
	},
	Layout: declarative.VBox{},
	Children: []declarative.Widget{
		declarative.Composite{
			Layout: declarative.VBox{},
			Children: []declarative.Widget{
				declarative.TextLabel{
					Text:      "The current fiber is not saved, do you really want to proceed ?",
					Alignment: declarative.Alignment2D(walk.AlignHCenterVCenter),
					Font:      declarative.Font{Family: "Roboto", PointSize: 9},
				},
				declarative.Composite{
					Layout: declarative.HBox{},
					Children: []declarative.Widget{
						declarative.PushButton{
							AssignTo: &acceptPB,
							Text:     "Yes",
							Font:     declarative.Font{Family: "Roboto", PointSize: 9},
							OnClicked: func() {
								Dlg.Accept()
							},
						},
						declarative.PushButton{
							AssignTo: &cancelPB,
							Text:     "Cancel",
							Font:     declarative.Font{Family: "Roboto", PointSize: 9},
							OnClicked: func() {
								Dlg.Cancel()
							},
						},
					},
				},
			},
		},
	},
}

// DeleteFiberButtonDialog This dialog is showed to confirm that user whant to delete a fiber's button
var DeleteFiberButtonDialog = declarative.Dialog{
	Icon:          "/icon.ico",
	Title:         "Confirm",
	AssignTo:      &Dlg,
	DefaultButton: &cancelPB,
	CancelButton:  &cancelPB,
	MinSize: declarative.Size{
		Width:  200,
		Height: 150,
	},
	Layout: declarative.VBox{},
	Children: []declarative.Widget{
		declarative.Composite{
			Layout: declarative.VBox{},
			Children: []declarative.Widget{
				declarative.TextLabel{
					Text:      "Do you really want to delete",
					Alignment: declarative.Alignment2D(walk.AlignHCenterVCenter),
					Font:      declarative.Font{Family: "Roboto", PointSize: 9},
				},
				declarative.Composite{
					Layout: declarative.HBox{},
					Children: []declarative.Widget{
						declarative.PushButton{
							AssignTo: &acceptPB,
							Text:     "YES",
							Font:     declarative.Font{Family: "Roboto", PointSize: 9},
							OnClicked: func() {
								Dlg.Accept()
							},
						},
						declarative.PushButton{
							AssignTo: &cancelPB,
							Text:     "NO",
							Font:     declarative.Font{Family: "Roboto", PointSize: 9},
							OnClicked: func() {
								Dlg.Cancel()
							},
						},
					},
				},
			},
		},
	},
}

// NoFiberNameSetDialog This dialog is showed when a user whant to save a fiber without having a name setteed
var NoFiberNameSetDialog = declarative.Dialog{
	Icon:          "/icon.ico",
	Title:         "Error",
	AssignTo:      &Dlg,
	DefaultButton: &acceptPB,
	MinSize: declarative.Size{
		Width:  200,
		Height: 150,
	},
	Layout: declarative.VBox{},
	Children: []declarative.Widget{
		declarative.Composite{
			Layout: declarative.VBox{},
			Children: []declarative.Widget{
				declarative.TextLabel{
					Text:      "No name given for the fiber",
					Alignment: declarative.Alignment2D(walk.AlignHCenterVCenter),
					Font:      declarative.Font{Family: "Roboto", PointSize: 9},
				},
				declarative.PushButton{
					AssignTo: &acceptPB,
					Text:     "OK",
					Font:     declarative.Font{Family: "Roboto", PointSize: 9},
					OnClicked: func() {
						Dlg.Cancel()
					},
				},
			},
		},
	},
}
