package log

import "fmt"

// Print log a value
func Print(msg interface{}, finished chan bool) {
	fmt.Println("FIBER INFO: Logging ...")
	fmt.Println("LOG: ", msg)
	finished <- true
}
