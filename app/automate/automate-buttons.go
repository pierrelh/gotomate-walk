package automate

import (
	"fmt"
	"gotomate/app/fiber"
	"gotomate/packages"
	"reflect"

	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
)

//NewButtons Create the automate's list of buttons
var NewButtons = new(Buttons)

//Buttons set the list of the fiber's buttons
type Buttons struct {
	Buttons []*Button
}

// Button Setting fiber's buttons structure
type Button struct {
	Composite *walk.Composite
	LinkLabel *walk.LinkLabel
	Dialog    *packages.Dialog
}

//CleanButtons Delete all the automate's buttons & delete all the button
func (btns *Buttons) CleanButtons() {
	for i := 0; i < len(btns.Buttons); i++ {
		btns.Buttons[i].Composite.Dispose()
		btns.Buttons[i].LinkLabel.Dispose()
	}
	p := reflect.ValueOf(btns).Elem()
	p.Set(reflect.Zero(p.Type()))
}

//DeleteButton Delete the button from automate's screen & remove the associated instruction from the fiber
func (btn *Button) DeleteButton(aw *Window, currentFiber *fiber.Fiber) {
	var dlg *walk.Dialog
	var acceptPB, cancelPB *walk.PushButton

	confirmDialog := declarative.Dialog{
		Icon:          "/icon.ico",
		Title:         "Confirm",
		AssignTo:      &dlg,
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
								AssignTo: &cancelPB,
								Text:     "NO",
								Font:     declarative.Font{Family: "Roboto", PointSize: 9},
								OnClicked: func() {
									dlg.Cancel()
								},
							},
							declarative.PushButton{
								AssignTo: &acceptPB,
								Text:     "YES",
								Font:     declarative.Font{Family: "Roboto", PointSize: 9},
								OnClicked: func() {
									for i := 0; i < len(NewButtons.Buttons); i++ {
										if btn == NewButtons.Buttons[i] {
											NewButtons.Buttons = append(NewButtons.Buttons[:i], NewButtons.Buttons[i+1:]...)
											currentFiber.Instructions = append(currentFiber.Instructions[:i], currentFiber.Instructions[i+1:]...)
											btn.Composite.Dispose()
											btn.LinkLabel.Dispose()
											dlg.Accept()
											return
										}
									}
									fmt.Println("ERROR: Unable to delete the fiber's instruction")
									dlg.Accept()
								},
							},
						},
					},
				},
			},
		},
	}
	confirmDialog.Run(aw.MainWindow)
}
