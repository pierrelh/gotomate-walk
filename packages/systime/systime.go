package systime

import (
	"fmt"
	"time"
)

// GetCurrentSysTime return the current sys time
func GetCurrentSysTime(finished chan bool) time.Time {
	fmt.Println("FIBER: Getting sys time...")
	t := time.Now()
	finished <- true
	return t
}

// GetCurrentSysClock return the current sys clock
func GetCurrentSysClock(finished chan bool) (int, int, int) {
	fmt.Println("FIBER: Getting sys clock...")
	h, m, s := time.Now().Clock()
	finished <- true
	return h, m, s
}
