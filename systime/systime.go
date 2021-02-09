package systime

import (
	"time"
)

// GetCurrentSysTime return the current sys time
func GetCurrentSysTime() time.Time {
	t := time.Now()
	return t
}

// GetCurrentSysClock return the current sys clock
func GetCurrentSysClock() (int, int, int) {
	h, m, s := time.Now().Clock()
	return h, m, s
}
