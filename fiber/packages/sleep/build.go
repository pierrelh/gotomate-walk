package sleep

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a sleep instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "Sleep":
		return new(Databinder), SleepTemplate
	case "MilliSleep":
		return new(Databinder), MilliSleepTemplate
	}
	fmt.Println("ERROR: Unable to find the function")
	return nil, nil
}
