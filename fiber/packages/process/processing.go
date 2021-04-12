package process

import (
	"fmt"
	"gotomate/fiber/value"
	"reflect"
	"regexp"
	"strconv"
)

// Processing process the functions from process's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) {
	switch funcName {
	case "StartProcess":
		go func() {
			path := instructionData.FieldByName("Path").Interface().(string)
			value.SetValue(instructionData.FieldByName("Output").Interface().(string), StartProcess(finished, path))
		}()
		<-finished
	case "KillProcess":
		Spid := instructionData.FieldByName("PID").Interface().(string)
		Ipid, _ := strconv.Atoi(Spid)
		if matched, _ := regexp.MatchString(`=.*`, Spid); matched {
			if val := value.KeySearch(Spid).Value; val != nil {
				Ipid = val.(int)
			} else {
				return
			}
		}
		go KillProcess(finished, Ipid)
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
}
