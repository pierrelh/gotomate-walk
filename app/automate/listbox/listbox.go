package listbox

import (
	"gotomate/app/automate/listbox/listboxitem"
	"gotomate/app/automate/listbox/listboxmodel"
	"gotomate/fiber/packages"
	"sort"
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
	var allPackages []string
	m := &listboxmodel.Model{Items: make([]listboxitem.Item, len(env))}

	for name := range env {
		allPackages = append(allPackages, name)
	}
	sort.Strings(allPackages)

	for i, name := range allPackages {
		value := strconv.Itoa(i)
		m.Items[i] = listboxitem.Item{Name: name, Value: value}
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
