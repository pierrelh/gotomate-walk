package process

import (
	"fmt"
	"os"
	"os/exec"
)

// StartProcess Create a process with a given path & return the process's pid
func StartProcess(finished chan bool, path string) int {
	fmt.Println("FIBER: Starting new process ...")
	cmd := exec.Command(path)
	cmd.Env = os.Environ()
	cmd.Start()
	finished <- true
	return cmd.Process.Pid
}
