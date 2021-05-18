package menu

import "github.com/lxn/walk"

//Menu Setting the automate window's menu structure
type Menu struct {
	FileMenu      *walk.Menu
	New           *walk.Action
	Import        *walk.Action
	Export        *walk.Action
	FiberMenu     *walk.Menu
	Save          *walk.Action
	Exit          *walk.Action
	FoldersMenu   *walk.Menu
	Run           *walk.Action
	Stop          *walk.Action
	Console       *walk.Action
	ImportPackage *walk.Action
	HelpMenu      *walk.Menu
	Documentation *walk.Action
	Packages      *walk.Action
}
