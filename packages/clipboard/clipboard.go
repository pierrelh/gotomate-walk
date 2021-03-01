// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package clipboard

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

// Read read string from clipboard
func Read(finished chan bool) (string, error) {
	fmt.Println("FIBER: Clipboard Reading ...")
	content, err := robotgo.ReadAll()
	finished <- true
	return content, err
}

// Write write string to clipboard
func Write(text string, finished chan bool) error {
	fmt.Println("FIBER: Clipboard Writing ...")
	err := robotgo.WriteAll(text)
	finished <- true
	return err
}
