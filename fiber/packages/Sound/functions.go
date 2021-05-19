package sound

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"

	"github.com/itchyny/volume-go"
)

// GetMuted Get the status of mute
func GetMuted(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting Mute status ...")

	mute, _ := volume.GetMuted()
	name := instructionData.FieldByName("Output").Interface().(string)
	variable.SetVariable(name, mute)
	finished <- true
	return -1
}

// GetVolume Get the current level of the volume
func GetVolume(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Get Volume ...")

	volume, _ := volume.GetVolume()
	name := instructionData.FieldByName("Output").Interface().(string)
	variable.SetVariable(name, volume)
	finished <- true
	return -1
}

// Mute the system's sound
func Mute(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Muting ...")
	volume.Mute()
	finished <- true
	return -1
}

// SetVolume Set the system's volume to the wanted value
func SetVolume(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Setting Volume ...")

	var volumeLevel int
	if volumeIsVar := instructionData.FieldByName("VolumeIsVar").Interface().(bool); volumeIsVar {
		volumeVarName := instructionData.FieldByName("VolumeVarName").Interface().(string)
		if val := variable.SearchVariable(volumeVarName).Value; val != nil {
			volumeLevel = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", volumeVarName)
			finished <- true
			return -1
		}
	} else {
		volumeLevel = instructionData.FieldByName("Volume").Interface().(int)
	}

	volume.SetVolume(volumeLevel)
	finished <- true
	return -1
}

// UnMute UnMute the system's sound
func UnMute(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: UnMuting ...")
	volume.Unmute()
	finished <- true
	return -1
}
