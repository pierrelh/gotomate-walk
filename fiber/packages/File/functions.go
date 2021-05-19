package file

import (
	"fmt"
	"gotomate/fiber/variable"
	"io/ioutil"
	"os"
	"reflect"
)

// Create a new file
func Create(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Creating File ...")

	var path string
	if IsVar := instructionData.FieldByName("PathIsVar").Interface().(bool); IsVar {
		pathVarName := instructionData.FieldByName("PathVarName").Interface().(string)
		if val := variable.SearchVariable(pathVarName).Value; val != nil {
			path = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", pathVarName)
			finished <- true
			return -1
		}
	} else {
		path = instructionData.FieldByName("Path").Interface().(string)
	}

	f, _ := os.Create(path)
	f.Close()
	finished <- true
	return -1
}

// Delete an existing file
func Delete(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Deleting File ...")

	var path string
	if IsVar := instructionData.FieldByName("PathIsVar").Interface().(bool); IsVar {
		pathVarName := instructionData.FieldByName("PathVarName").Interface().(string)
		if val := variable.SearchVariable(pathVarName).Value; val != nil {
			path = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", pathVarName)
			finished <- true
			return -1
		}
	} else {
		path = instructionData.FieldByName("Path").Interface().(string)
	}

	os.Remove(path)
	finished <- true
	return -1
}

// Read a file
func Read(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Reading File ...")

	var path string
	if IsVar := instructionData.FieldByName("PathIsVar").Interface().(bool); IsVar {
		pathVarName := instructionData.FieldByName("PathVarName").Interface().(string)
		if val := variable.SearchVariable(pathVarName).Value; val != nil {
			path = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", pathVarName)
			finished <- true
			return -1
		}
	} else {
		path = instructionData.FieldByName("Path").Interface().(string)
	}

	content, _ := ioutil.ReadFile(path)
	name := instructionData.FieldByName("Output").Interface().(string)
	variable.SetVariable(name, string(content))
	finished <- true
	return -1
}

// Write an existing file
func Write(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Reading File ...")

	var path string
	if pathIsVar := instructionData.FieldByName("PathIsVar").Interface().(bool); pathIsVar {
		pathVarName := instructionData.FieldByName("PathVarName").Interface().(string)
		if val := variable.SearchVariable(pathVarName).Value; val != nil {
			path = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", pathVarName)
			finished <- true
			return -1
		}
	} else {
		path = instructionData.FieldByName("Path").Interface().(string)
	}

	var content string
	if contentIsVar := instructionData.FieldByName("ContentIsVar").Interface().(bool); contentIsVar {
		ContentVarName := instructionData.FieldByName("ContentVarName").Interface().(string)
		if val := variable.SearchVariable(ContentVarName).Value; val != nil {
			content = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", ContentVarName)
			finished <- true
			return -1
		}
	} else {
		content = instructionData.FieldByName("Content").Interface().(string)
	}

	ioutil.WriteFile(path, []byte(content), 0644)
	finished <- true
	return -1
}
