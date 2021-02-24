package log

import "fmt"

// Print log a value
func Print(msg string, finished chan bool) {
	fmt.Println("Log initialization ...")
	fmt.Println(msg)
	finished <- true
}
