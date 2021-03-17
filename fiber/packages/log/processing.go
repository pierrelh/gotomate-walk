package log

import (
	"fmt"
	"gotomate/fiber/value"
	"reflect"
	"regexp"
	"strings"
)

// Processing process the functions from log's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) {
	switch funcName {
	case "Print":
		logValue := instructionData.FieldByName("Log").Interface().(string)
		matched, _ := regexp.MatchString(`=.*`, logValue)
		msg := instructionData.FieldByName("Log").Interface()
		if matched {
			varSlice := strings.Split(logValue, "=")
			for i := 0; i < len(varSlice); i++ {
				if varSlice[i] != "" {
					varInstruction := value.KeySearch(regexp.QuoteMeta(varSlice[i]))
					if varInstruction != nil {
						msg = reflect.ValueOf(varInstruction.Value)
					} else {
						fmt.Println("FIBER ERROR: Unable to find the fiber's var: ", regexp.QuoteMeta(varSlice[i]))
						msg = ""
					}
				}
			}
		}
		go Print(msg, finished)
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
}
