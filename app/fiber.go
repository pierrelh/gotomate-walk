package app

import (
	"fmt"
	"gotomate/sleep"
	"reflect"
	"time"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

// Fiber Initialize the fiber structure
type Fiber struct {
	Instructions []*FiberInstruction
}

// FiberInstruction Initialize a fiber's instruction
type FiberInstruction struct {
	Package  string
	FuncName string
	Button   *FiberButton
}

// FiberButton Setting fiber's buttons structure
type FiberButton struct {
	*walk.Composite
	*walk.ImageView
	*walk.LinkLabel
	DialogWindow Dialog
}

// FiberComposite Setting FiberButton's composite & it's dialog
type FiberComposite struct {
	*walk.Composite
	DialogWindow Dialog
}

// FiberImageView Setting FiberButton's imageview & it's dialog
type FiberImageView struct {
	*walk.ImageView
	DialogWindow Dialog
}

// FiberLinkLabel Setting FiberButton's linklabel & it's dialog
type FiberLinkLabel struct {
	*walk.LinkLabel
	DialogWindow Dialog
}

//WndProc setting the window event of the composite element
func (fb *FiberComposite) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_LBUTTONDOWN:
		go fb.DialogWindow.Run(aw.mw)
	}

	return fb.Composite.WndProc(hwnd, msg, wParam, lParam)
}

//WndProc setting the window event of the imageview element
func (fb *FiberImageView) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_LBUTTONDOWN:
		go fb.DialogWindow.Run(aw.mw)
	}

	return fb.ImageView.WndProc(hwnd, msg, wParam, lParam)
}

//WndProc setting the window event of the linklabel element
func (fb *FiberLinkLabel) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_LBUTTONDOWN:
		fmt.Println("left down")
		go fb.DialogWindow.Run(aw.mw)
	}

	return fb.LinkLabel.WndProc(hwnd, msg, wParam, lParam)
}

func runFiber() {
	fmt.Println("| Fiber Start |")

	for i := 0; i < len(fiber.Instructions); i++ {
		finished := make(chan bool)
		instruction := fiber.Instructions[i]
		switch instruction.Package {
		case "Sleep":
			switch instruction.FuncName {
			case "Sleep":
				val := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
				duration := val.FieldByName("Duration").Interface().(time.Duration)
				go sleep.Sleep(duration, finished)
				<-finished
			}
		}
	}
	fmt.Println("| Fiber Finished |")
}
