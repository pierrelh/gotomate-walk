package app

import (
	"gotomate/log"

	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
)

// CreateNewDialog Create the dialog for the function & add the needed template
func CreateNewDialog(funcName string) declarative.Dialog {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton
	source, children := fillDialog(funcName)

	return declarative.Dialog{
		Icon:          "/icon.ico",
		Title:         "Settings",
		AssignTo:      &dlg,
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: declarative.DataBinder{
			AssignTo:       &db,
			Name:           funcName,
			DataSource:     source,
			ErrorPresenter: declarative.ToolTipErrorPresenter{},
		},
		MinSize: declarative.Size{
			Width:  300,
			Height: 300,
		},
		Layout: declarative.VBox{},
		Children: []declarative.Widget{
			declarative.Composite{
				Layout:   declarative.VBox{},
				Children: children,
			},
			declarative.HSplitter{
				MaxSize: declarative.Size{Height: 30},
				Children: []declarative.Widget{
					declarative.HSpacer{},
					declarative.PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						Font:     declarative.Font{Family: "Segoe UI", PointSize: 9},
						OnClicked: func() {
							if err := db.Submit(); err != nil {
								log.Print(err)
								return
							}

							dlg.Accept()
						},
					},
					declarative.PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						Font:      declarative.Font{Family: "Segoe UI", PointSize: 9},
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}
}

func fillDialog(funcName string) (interface{}, []declarative.Widget) {
	switch funcName {
	case "Sleep":
		return new(Sleep), SleepTemplate
	case "MilliSleep":
		return new(MilliSleep), MilliSleepTemplate
	case "Click":
		return new(MouseClick), MouseClickTemplate
	case "Scroll":
		return new(MouseScroll), MouseScrollTemplate
	case "Move":
		return new(MouseMove), MouseMoveTemplate
	case "Tap":
		return new(KeyboardTap), KeyboardTapTemplate
	case "Write":
		return new(ClipboardWrite), ClipboardWriteTemplate
	default:
		return nil, nil
	}
}
