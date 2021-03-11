package notification

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a notification instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "Create":
		return new(CreateDatabinder), CreateTemplate
	}
	fmt.Println("ERROR: Unable to find the function")
	return nil, nil
}
