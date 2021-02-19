package main

import (
	"gotomate/app"

	"github.com/lxn/walk"
)

func main() {
	walk.Resources.SetRootDirPath("img")
	app.CreateApp()
}
