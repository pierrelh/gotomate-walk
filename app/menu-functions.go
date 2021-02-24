package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lxn/walk"
)

//AddSavedFibersActions Add all the saved fibers to the My fibers's menu
func AddSavedFibersActions() {
	root := "./saves"
	aw.folders.Actions().Clear()

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			var file = filepath.Base(path)
			var extension = filepath.Ext(file)
			var name = file[0 : len(file)-len(extension)]
			a := walk.NewAction()
			a.SetText(name)
			a.Triggered().Attach(func() { OpenSavedFiber(name) })
			aw.folders.Actions().Add(a)
		}
		return nil
	})
}

//OpenSavedFiber Open a saved fiber from the menu My Fibers
func OpenSavedFiber(name string) {
	fmt.Println(name)
}
