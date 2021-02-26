package menu

import "github.com/lxn/walk"

//Menu Setting the automate window's menu structure
type Menu struct {
	WindowMenu *walk.Menu
	New        *walk.Action
	Open       *walk.Action
	Folders    *walk.Menu
	Run        *walk.Action
	Stop       *walk.Action
	Save       *walk.Action
	Exit       *walk.Action
}
