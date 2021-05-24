package process

import (
	"fmt"
	"gotomate/fiber/variable"
	"os"
	"os/exec"
	"reflect"

	"github.com/go-vgo/robotgo"
)

// GetTitle get the title of a process
func GetTitle(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting title of process ...")

	var pid int
	if pidIsVar := instructionData.FieldByName("PIDIsVar").Interface().(bool); pidIsVar {
		PIDVarName := instructionData.FieldByName("PIDVarName").Interface().(string)
		if val := variable.SearchVariable(PIDVarName).Value; val != nil {
			pid = val.(int)
		} else {
			finished <- true
			return -1
		}
	} else {
		pid = instructionData.FieldByName("PID").Interface().(int)
	}

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), robotgo.GetTitle(int32(pid)))
	finished <- true
	return -1
}

// GetPid get the pid of the active window
func GetPid(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting process pid ...")

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), robotgo.GetPID())
	finished <- true
	return -1
}

// KillProcess Kill a process by his pid
func KillProcess(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Killing process ...")

	var pid int
	if pidIsVar := instructionData.FieldByName("PIDIsVar").Interface().(bool); pidIsVar {
		PIDVarName := instructionData.FieldByName("PIDVarName").Interface().(string)
		if val := variable.SearchVariable(PIDVarName).Value; val != nil {
			pid = val.(int)
		} else {
			finished <- true
			return -1
		}
	} else {
		pid = instructionData.FieldByName("PID").Interface().(int)
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

// MaxSize set the max size for a process
func MaxSize(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Maximizing process size ...")

	var pid int
	if pidIsVar := instructionData.FieldByName("PIDIsVar").Interface().(bool); pidIsVar {
		PIDVarName := instructionData.FieldByName("PIDVarName").Interface().(string)
		if val := variable.SearchVariable(PIDVarName).Value; val != nil {
			pid = val.(int)
		} else {
			finished <- true
			return -1
		}
	} else {
		pid = instructionData.FieldByName("PID").Interface().(int)
	}

	robotgo.MaxWindow(int32(pid))
	finished <- true
	return -1
}

// Reduce a process
func Reduce(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Reducing a process ...")

	var pid int
	if pidIsVar := instructionData.FieldByName("PIDIsVar").Interface().(bool); pidIsVar {
		PIDVarName := instructionData.FieldByName("PIDVarName").Interface().(string)
		if val := variable.SearchVariable(PIDVarName).Value; val != nil {
			pid = val.(int)
		} else {
			finished <- true
			return -1
		}
	} else {
		pid = instructionData.FieldByName("PID").Interface().(int)
	}

	robotgo.MinWindow(int32(pid))
	finished <- true
	return -1
}

// StartProcess Create a process with a given path & return the process's pid
func StartProcess(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	var path string
	if pathIsVar := instructionData.FieldByName("PathIsVar").Interface().(bool); pathIsVar {
		pathVarName := instructionData.FieldByName("PathVarName").Interface().(string)
		if val := variable.SearchVariable(pathVarName).Value; val != nil {
			path = val.(string)
		} else {
			finished <- true
			return -1
		}
	} else {
		path = instructionData.FieldByName("Path").Interface().(string)
	}

	cmd := exec.Command(path)
	cmd.Env = os.Environ()
	cmd.Start()
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), cmd.Process.Pid)
	finished <- true
	return -1
}
