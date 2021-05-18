package process

// StartProcessDatabinder Define the StartProcess databinder
type StartProcessDatabinder struct {
	Path        string
	PathVarName string
	PathIsVar   bool
	Output      string
}

// KillProcessDatabinder Define the KillProcess databinder
type KillProcessDatabinder struct {
	PID        int
	PIDVarName string
	PIDIsVar   bool
}
