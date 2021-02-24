package app

import (
	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
)

// FiberButton Setting fiber's buttons structure
type FiberButton struct {
	Composite    *walk.Composite
	LinkLabel    *walk.LinkLabel
	DialogWindow declarative.Dialog
}

// CreateFiberButton Create a new Fiberbutton in the fiber
func CreateFiberButton(parent *walk.ScrollView, funcName, packageName string, dialog declarative.Dialog) error {
	fb := new(FiberButton)
	fb.DialogWindow = dialog
	bmp, err := walk.NewBitmapFromFile(walk.Resources.RootDirPath() + "/func-icons/" + packageName + ".png")
	if err != nil {
		bmp, _ = walk.NewBitmapFromFile(walk.Resources.RootDirPath() + "/func-icons/default.png")
	}

	if err := (declarative.Composite{
		AssignTo:   &fb.Composite,
		Layout:     declarative.VBox{MarginsZero: true, SpacingZero: true},
		Background: declarative.BitmapBrush{Image: bmp},
		Alignment:  declarative.Alignment2D(walk.AlignHNearVCenter),
		OnMouseDown: func(x, y int, button walk.MouseButton) {
			fb.DialogWindow.Run(aw.MainWindow)
		},
		Children: []declarative.Widget{
			declarative.HSpacer{},
			declarative.LinkLabel{
				AssignTo:  &fb.LinkLabel,
				Font:      declarative.Font{Family: "Roboto", PointSize: 9, Bold: true},
				Text:      funcName,
				Alignment: declarative.Alignment2D(walk.AlignHCenterVCenter),
			},
			declarative.HSpacer{},
		},
	}).Create(declarative.NewBuilder(parent)); err != nil {
		return err
	}
	fb.Composite.SetCursor(walk.CursorHand())
	fb.Composite.SetMinMaxSizePixels(walk.Size{Width: 300, Height: 150}, walk.Size{Width: 300, Height: 150})
	fb.LinkLabel.SetMinMaxSizePixels(walk.Size{Width: 300, Height: 20}, walk.Size{Width: 300, Height: 20})
	fb.Composite.Children().At(0).SetMinMaxSizePixels(walk.Size{Width: 300, Height: 100}, walk.Size{Width: 300, Height: 100})
	fb.Composite.Children().At(2).SetMinMaxSizePixels(walk.Size{Width: 300, Height: 30}, walk.Size{Width: 300, Height: 30})
	return nil
}
