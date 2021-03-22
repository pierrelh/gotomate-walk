package listbox

import (
	"fmt"
	"gotomate/app/window/listbox/listboxitem"
	"gotomate/app/window/listbox/listboxmodel"
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
	m := &listboxmodel.Model{Items: make([]listboxitem.Item, len(env))}

	for i := 0; i < len(env); i++ {
		value := strconv.Itoa(i)
		m.Items[i] = listboxitem.Item{Name: env[i].Name, Value: value}
	}
	return m
}

//NewSecondaryListModel Fill the SecondaryListModel with the adapted datas
func (list *ListBox) NewSecondaryListModel(key string) *listboxmodel.Model {
	env := packages.Packages
	idx := sort.Search(len(env), func(i int) bool {
		return env[i].Name >= key
	})
	if idx == len(env) {
		fmt.Println("GOTOMATE ERROR: Unable to find", key, "package")
		return nil
	} else {
		m := &listboxmodel.Model{Items: make([]listboxitem.Item, len(env[idx].Functions))}

		for i := 0; i < len(env[idx].Functions); i++ {
			value := strconv.Itoa(i)
			m.Items[i] = listboxitem.Item{Name: env[idx].Functions[i], Value: value}
		}
		return m
	}
}
