package app

import (
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
