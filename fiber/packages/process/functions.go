package process

import (
	"fmt"
	"gotomate/fiber/value"
	"os"
	"os/exec"
	"reflect"
	"strconv"
)

func KillProcess(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Killing process ...")
	Spid := instructionData.FieldByName("PID").Interface().(string)
	pid, _ := strconv.Atoi(Spid)
	if pidIsVar := instructionData.FieldByName("PIDIsVar").Interface().(bool); pidIsVar {
		if val := value.KeySearch(Spid).Value; val != nil {
			pid = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", Spid)
			finished <- true
			return -1
		}
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("FIBER WARNING: Didn't find process with pid", pid)
		finished <- true
		return -1
	}
	proc.Kill()
	finished <- true
	return -1
}

// StartProcess Create a process with a given path & return the process's pid
func StartProcess(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")
	path := instructionData.FieldByName("Path").Interface().(string)
	if pathIsVar := instructionData.FieldByName("PathIsVar").Interface().(bool); pathIsVar {
		if val := value.KeySearch(path).Value; val != nil {
			path = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", path)
			finished <- true
			return -1
		}
	}

	cmd := exec.Command(path)
	cmd.Env = os.Environ()
	cmd.Start()
	value.SetValue(instructionData.FieldByName("Output").Interface().(string), cmd.Process.Pid)
	finished <- true
	return -1
}
