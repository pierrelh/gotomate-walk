package app

import (
	"fmt"
	"gotomate/clipboard"
	"gotomate/keyboard"
	"gotomate/mouse"
	"gotomate/sleep"
	"reflect"

	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
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
	DialogWindow declarative.Dialog
}

// FiberComposite Setting FiberButton's composite & it's dialog
type FiberComposite struct {
	*walk.Composite
	DialogWindow declarative.Dialog
}

// FiberImageView Setting FiberButton's imageview & it's dialog
type FiberImageView struct {
	*walk.ImageView
	DialogWindow declarative.Dialog
}

// FiberLinkLabel Setting FiberButton's linklabel & it's dialog
type FiberLinkLabel struct {
	*walk.LinkLabel
	DialogWindow declarative.Dialog
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
				duration := val.FieldByName("Duration").Interface().(float64)
				go sleep.Sleep(duration, finished)
				<-finished
			case "MilliSleep":
				val := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
				duration := val.FieldByName("Duration").Interface().(float64)
				go sleep.MilliSleep(duration, finished)
				<-finished
			default:
				fmt.Println("This function is not integrated yet: " + instruction.FuncName)
				continue
			}
		case "Mouse":
			switch instruction.FuncName {
			case "Click":
				val := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
				button := val.FieldByName("MouseButtonName").Interface().(string)
				go mouse.Click(button, finished)
				<-finished
			case "Scroll":
				val := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
				x := val.FieldByName("X").Interface().(int)
				y := val.FieldByName("Y").Interface().(int)
				go mouse.Scroll(x, y, finished)
				<-finished
			case "Move":
				val := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
				x := val.FieldByName("X").Interface().(int)
				y := val.FieldByName("Y").Interface().(int)
				go mouse.Move(x, y, finished)
				<-finished
			default:
				fmt.Println("This function is not integrated yet: " + instruction.FuncName)
				continue
			}
		case "Keyboard":
			switch instruction.FuncName {
			case "Tap":
				val := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
				input := val.FieldByName("Input").Interface().(string)
				special := []string{}
				if err := val.FieldByName("Special1").Interface().(string); err != "" {
					special = append(special, err)
				}
				if err := val.FieldByName("Special2").Interface().(string); err != "" {
					special = append(special, err)
				}
				go keyboard.Tap(input, special, finished)
				<-finished
			default:
				fmt.Println("This function is not integrated yet: " + instruction.FuncName)
				continue
			}
		case "Clipboard":
			switch instruction.FuncName {
			case "Write":
				val := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
				content := val.FieldByName("Content").Interface().(string)
				go clipboard.Write(content, finished)
				<-finished
			case "Read":
				go func() {
					content, _ := clipboard.Read(finished)
					val := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
					val.FieldByName("Content").SetString(content)
					fmt.Println(val.FieldByName("Content"))
				}()
				<-finished
			default:
				fmt.Println("This function is not integrated yet: " + instruction.FuncName)
				continue
			}
		default:
			fmt.Println("This package is not integrated yet: " + instruction.Package)
			continue
		}
	}
	fmt.Println("| Fiber Finished |")
}
