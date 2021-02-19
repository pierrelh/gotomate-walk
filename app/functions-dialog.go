package app

import (
	"gotomate/log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// CreateNewDialog Create the dialog for the function & add the needed template
func CreateNewDialog(funcName string) Dialog {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton
	source, children := fillDialog(funcName)

	return Dialog{
		Icon:          "/icon.ico",
		Title:         "Settings",
		AssignTo:      &dlg,
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			Name:           funcName,
			DataSource:     source,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		MinSize: Size{
			Width:  300,
			Height: 300,
		},
		Layout: VBox{},
		Children: []Widget{
			Composite{
				Layout:   Grid{Columns: 2},
				Children: children,
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							if err := db.Submit(); err != nil {
								log.Print(err)
								return
							}

							dlg.Accept()
						},
					},
					PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}
}

func fillDialog(funcName string) (interface{}, []Widget) {
	switch funcName {
	case "Sleep":
		return new(Sleep), SleepTemplate
	case "MilliSleep":
		return new(MilliSleep), MilliSleepTemplate
	default:
		return nil, nil
	}
}
