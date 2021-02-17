package app

import (
	"gotomate/log"
	"time"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

//Sleep Define the Sleep parameters
type Sleep struct {
	Duration time.Duration
}

type DurationField struct {
	p *time.Duration
}

func (s *Sleep) SleepDurationField() *DurationField {
	return &DurationField{&s.Duration}
}

// CreateNewDialog Create the dialog for the function & add the needed template
func CreateNewDialog(funcName string) Dialog {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		Icon:          "img/icon.ico",
		Title:         "Settings",
		AssignTo:      &dlg,
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			Name:           "sleep",
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		MinSize: Size{300, 300},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout:   Grid{Columns: 2},
				Children: fillDialog(funcName),
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

func fillDialog(funcName string) []Widget {
	switch funcName {
	case "Sleep":
		return SleepTemplate
	default:
		return nil

	}
}
