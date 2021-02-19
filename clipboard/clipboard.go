// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package clipboard

import (
	"fmt"
	"syscall"
	"time"

	"github.com/go-vgo/robotgo"
)

const (
	cfUnicodetext = 13
	// gmemFixed     = 0x0000
	gmemMoveable = 0x0002
)

var (
	user32           = syscall.MustLoadDLL("user32")
	openClipboard    = user32.MustFindProc("OpenClipboard")
	closeClipboard   = user32.MustFindProc("CloseClipboard")
	emptyClipboard   = user32.MustFindProc("EmptyClipboard")
	getClipboardData = user32.MustFindProc("GetClipboardData")
	setClipboardData = user32.MustFindProc("SetClipboardData")

	kernel32     = syscall.NewLazyDLL("kernel32")
	globalAlloc  = kernel32.NewProc("GlobalAlloc")
	globalFree   = kernel32.NewProc("GlobalFree")
	globalLock   = kernel32.NewProc("GlobalLock")
	globalUnlock = kernel32.NewProc("GlobalUnlock")
	lstrcpy      = kernel32.NewProc("lstrcpyW")
)

// waitOpenClipboard opens the clipboard, waiting for up to a second to do so.
func waitOpenClipboard() error {
	started := time.Now()
	limit := started.Add(time.Second)
	var (
		r   uintptr
		err error
	)
	for time.Now().Before(limit) {
		r, _, err = openClipboard.Call(0)
		if r != 0 {
			return nil
		}
		time.Sleep(time.Millisecond)
	}

	return err
}

// Read read string from clipboard
func Read(finished chan bool) (string, error) {
	fmt.Println("Clipboard Reading initialization ...")
	content, err := robotgo.ReadAll()
	finished <- true
	return content, err
}

// Write write string to clipboard
func Write(text string, finished chan bool) error {
	fmt.Println("Clipboard Writing initialization ...")
	err := robotgo.WriteAll(text)
	finished <- true
	return err
}
