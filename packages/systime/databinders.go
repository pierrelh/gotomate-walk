package systime

import "time"

// ClockDatabinder Define the SysClock parameters
type ClockDatabinder struct {
	Var   string
	Value [3]int
}

// TimeDatabinder Define the SysTime parameters
type TimeDatabinder struct {
	Var   string
	Value time.Time
}
