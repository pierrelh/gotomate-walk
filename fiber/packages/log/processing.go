package log

import (
	"fmt"
	"gotomate/fiber/value"
	"reflect"
	"regexp"
)

// Processing process the functions from log's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) {
	switch funcName {
	case "Print":
		logValue := instructionData.FieldByName("Log").Interface()
		if matched, _ := regexp.MatchString(`=.*`, logValue.(string)); matched {
			logValue = value.KeySearch(logValue.(string)).Value
		}
		go Print(logValue, finished)
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
}
