package notification

import (
	"fmt"
	"gotomate/fiber/value"
	"reflect"
	"regexp"
)

// Processing process the functions from notification's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) {
	switch funcName {
	case "Create":
		title := instructionData.FieldByName("Title").Interface().(string)
		if matched, _ := regexp.MatchString(`=.*`, title); matched {
			if val := value.KeySearch(title).Value; val != nil {
				title = val.(string)
			} else {
				return
			}
		}
		msg := instructionData.FieldByName("Message").Interface().(string)
		if matched, _ := regexp.MatchString(`=.*`, msg); matched {
			if val := value.KeySearch(msg).Value; val != nil {
				msg = val.(string)
			} else {
				return
			}
		}
		go Create(title, msg, finished)
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
}
