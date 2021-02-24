package fiber

import (
	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
)

// Button Setting fiber's buttons structure
type Button struct {
	Composite    *walk.Composite
	LinkLabel    *walk.LinkLabel
	DialogWindow declarative.Dialog
}
