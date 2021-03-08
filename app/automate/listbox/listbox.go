package listbox

import (
	"gotomate/app/automate/listbox/listboxitem"
	"gotomate/app/automate/listbox/listboxmodel"
	"gotomate/fiber/packages"
	"strconv"

	"github.com/lxn/walk"
)

//ListBox define the structure of a automate List
type ListBox struct {
	ListBox *walk.ListBox
	Model   *listboxmodel.Model
}

// NewPrimaryListModel Getting the automates packages
func NewPrimaryListModel() *listboxmodel.Model {
	env := packages.Packages
	i := 0

	m := &listboxmodel.Model{Items: make([]listboxitem.Item, len(env))}

	for name := range env {
		value := strconv.Itoa(i)

		m.Items[i] = listboxitem.Item{Name: name, Value: value}
		i++
	}

	return m
}

//NewSecondaryListModel Fill the SecondaryListModel with the adapted datas
func (list *ListBox) NewSecondaryListModel(key string) *listboxmodel.Model {
	env := packages.Packages

	m := &listboxmodel.Model{Items: make([]listboxitem.Item, len(env[key]))}

	for i, e := range env[key] {
		value := strconv.Itoa(i)
		name := e

		m.Items[i] = listboxitem.Item{Name: name, Value: value}
	}

	return m
}
