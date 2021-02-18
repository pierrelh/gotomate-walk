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
	*walk.PushButton
	DialogWindow Dialog
}

// WndProc Trigger the button's events
func (fb *FiberButton) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_LBUTTONDOWN:
		go fb.DialogWindow.Run(aw.mw)
	}

	return fb.PushButton.WndProc(hwnd, msg, wParam, lParam)
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
