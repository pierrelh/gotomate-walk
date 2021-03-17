package sleep

import (
	"fmt"
	"time"
)

// MilliSleep sleep tm milli second
func MilliSleep(tm float64, finished chan bool) {
	fmt.Println("FIBER INFO: Sleeping for: ", tm, "ms")
	time.Sleep(time.Duration(tm) * time.Millisecond)
	finished <- true
}

// Sleep time.Sleep tm second
func Sleep(tm float64, finished chan bool) {
	fmt.Println("FIBER INFO: Sleeping for: ", tm, "s")
	time.Sleep(time.Duration(tm) * time.Second)
	finished <- true
}
