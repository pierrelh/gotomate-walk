package process

import (
	"fmt"
	"os"
	"os/exec"
)

// StartProcess Create a process with a given path & return the process's pid
func StartProcess(finished chan bool, path string) int {
	fmt.Println("FIBER INFO: Starting new process ...")
	cmd := exec.Command(path)
	cmd.Env = os.Environ()
	cmd.Start()
	finished <- true
	return cmd.Process.Pid
}

func KillProcess(finished chan bool, pid int) {
	fmt.Println("FIBER INFO: Killing process ...")
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("FIBER WARNING: Didn't find process with pid", pid)
	}
	proc.Kill()
	finished <- true
}
