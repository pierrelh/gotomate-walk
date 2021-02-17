package app

import (
	"gotomate/log"
	"time"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type Sleep struct {
	Duration time.Duration
}

type DurationField struct {
	p *time.Duration
}

func (s *Sleep) SleepDurationField() *DurationField {
	return &DurationField{&s.Duration}
}

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
	widget := []Widget{}
	switch funcName {
	case "Sleep":
		widget = append(widget,
			Label{
				ColumnSpan: 2,
				Text:       "Sleep For:",
			},
			LineEdit{
				ColumnSpan: 2,
				Text:       Bind("SleepDurationField"),
			},
		)
	default:

	}
	return widget
}

// func SleepDialog(aw walk.Form) (int, error) {
// 	var dlg *walk.Dialog
// 	var db *walk.DataBinder
// 	var acceptPB, cancelPB *walk.PushButton

// 	return Dialog{
// 		Icon:          "img/icon.ico",
// 		Title:         "Settings",
// 		AssignTo:      &dlg,
// 		DefaultButton: &acceptPB,
// 		CancelButton:  &cancelPB,
// 		DataBinder: DataBinder{
// 			AssignTo:       &db,
// 			Name:           "sleep",
// 			ErrorPresenter: ToolTipErrorPresenter{},
// 		},
// 		MinSize: Size{300, 300},
// 		Layout:  VBox{},
// 		Children: []Widget{
// 			Composite{
// 				Layout: Grid{Columns: 2},
// 				Children: []Widget{
// 					Label{
// 						ColumnSpan: 2,
// 						Text:       "Sleep For:",
// 					},
// 					LineEdit{
// 						ColumnSpan: 2,
// 						Text:       Bind("SleepDurationField"),
// 					},
// 				},
// 			},
// 			Composite{
// 				Layout: HBox{},
// 				Children: []Widget{
// 					HSpacer{},
// 					PushButton{
// 						AssignTo: &acceptPB,
// 						Text:     "OK",
// 						OnClicked: func() {
// 							if err := db.Submit(); err != nil {
// 								log.Print(err)
// 								return
// 							}

// 							dlg.Accept()
// 						},
// 					},
// 					PushButton{
// 						AssignTo:  &cancelPB,
// 						Text:      "Cancel",
// 						OnClicked: func() { dlg.Cancel() },
// 					},
// 				},
// 			},
// 		},
// 	}.Run(aw)
// }
