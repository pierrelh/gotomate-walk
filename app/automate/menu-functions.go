package automate

import (
	"encoding/json"
	"fmt"
	"gotomate/app/fiber"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
)

//AddSavedFibersActions Add all the saved fibers to the My fibers's menu
func (aw *Window) AddSavedFibersActions() {
	root := "./saves"
	aw.Menu.Folders.Actions().Clear()

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			var file = filepath.Base(path)
			var extension = filepath.Ext(file)
			var name = file[0 : len(file)-len(extension)]
			a := walk.NewAction()
			a.SetText(name)
			a.Triggered().Attach(func() { aw.OpenSavedFiber(name) })
			aw.Menu.Folders.Actions().Add(a)
		}
		return nil
	})
}

//OpenSavedFiber Open a saved fiber from the menu My Fibers
func (aw *Window) OpenSavedFiber(name string) {
	fmt.Println(name)
}

//SaveFiber save the current fiber
func (aw *Window) SaveFiber(fiber *fiber.Fiber) {
	name := aw.FiberNameInput.Text()
	if name == "" {
		var dlg *walk.Dialog
		var acceptPB *walk.PushButton

		errDialog := declarative.Dialog{
			Icon:          "/icon.ico",
			Title:         "Error",
			AssignTo:      &dlg,
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
								dlg.Cancel()
								aw.FiberNameInput.SetFocus()
							},
						},
					},
				},
			},
		}

		errDialog.Run(aw.MainWindow)
	} else {
		fiber.Name = name
		path := "saves/" + name + ".json"
		file, _ := json.Marshal(fiber)
		ioutil.WriteFile(path, file, 0644)
		aw.AddSavedFibersActions()
	}
}
