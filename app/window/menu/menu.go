package menu

import "github.com/lxn/walk"

//Menu Setting the automate window's menu structure
type Menu struct {
	FileMenu      *walk.Menu
	New           *walk.Action
	Import        *walk.Action
	FoldersMenu   *walk.Menu
	Console       *walk.Action
	Save          *walk.Action
	Export        *walk.Action
	Exit          *walk.Action
	FiberMenu     *walk.Menu
	Run           *walk.Action
	Stop          *walk.Action
	HelpMenu      *walk.Menu
	Documentation *walk.Action
	Packages      *walk.Action
}
