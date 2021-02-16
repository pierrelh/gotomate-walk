package automate

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

func SleepDialog(aw walk.Form) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
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
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						ColumnSpan: 2,
						Text:       "Sleep For:",
					},
					LineEdit{
						ColumnSpan: 2,
						Text:       Bind("SleepDurationField"),
					},
				},
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
	}.Run(aw)
}

type Dialog struct {
	AssignTo      **walk.MainWindow
	Title         string
	DefaultButton **walk.PushButton
	CancelButton  **walk.PushButton
	DataBinder    []DataBinder
	MinSize       walk.Size
	Layout        walk.Layout
	Children      []walk.Widget
}

type DataBinder struct {
	AssignTo       **walk.DataBinder
	Name           string
	DataSource     interface{}
	ErrorPresenter walk.ErrorPresenter
}

type sleepDialog struct {
	Dialog []Dialog
}
