package button

import (
	"gotomate/fiber/packages"
	"reflect"

	"github.com/lxn/walk"
)

//NewButtons Create the automate's list of buttons
var NewButtons = new(Buttons)

//Buttons set the list of the fiber's buttons
type Buttons struct {
	Buttons []*Button
}

// Button Setting fiber's buttons structure
type Button struct {
	Composite   *walk.Composite
	IDLabel     *walk.LinkLabel
	FuncLabel   *walk.LinkLabel
	NextID      *walk.NumberEdit
	NextIDLabel *walk.Label
	Dialog      *packages.Dialog
}

//CleanButtons Delete all the automate's buttons & delete all the button
func (btns *Buttons) CleanButtons() {
	p := reflect.ValueOf(btns).Elem()
	p.Set(reflect.Zero(p.Type()))
}
