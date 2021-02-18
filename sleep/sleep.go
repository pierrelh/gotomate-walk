package sleep

import (
	"fmt"
	"time"
)

// MilliSleep sleep tm milli second
func MilliSleep(tm time.Duration) {
	fmt.Println("Sleeping for: ", tm)
	time.Sleep(tm * time.Millisecond)
}

// Sleep time.Sleep tm second
func Sleep(tm time.Duration, finished chan bool) {
	fmt.Println("Sleeping for: ", tm)
	time.Sleep(tm)
	finished <- true
}
