package listboxmodel

import (
	"gotomate/app/window/listbox/listboxitem"

	"github.com/lxn/walk"
)

//Model Setting the model of automate ListBox
type Model struct {
	walk.ListModelBase
	Items []listboxitem.Item
}

// ItemCount return the length of an AutomateModel items
func (m *Model) ItemCount() int {
	return len(m.Items)
}

// Value return the value of an AutomateModel item
func (m *Model) Value(index int) interface{} {
	return m.Items[index].Name
}
