package battery

import (
	"fmt"
	"gotomate/fiber/value"
	"reflect"

	batType "github.com/distatus/battery"
)

// Processing process the functions from battery's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) {
	switch funcName {
	case "GetBattery":
		go func() {
			value.SetValue(instructionData.FieldByName("Output").Interface().(string), GetBattery(finished))
		}()
		<-finished
	case "GetBatteryState":
		go func() {
			batName := instructionData.FieldByName("BatteryName").Interface().(string)
			if batInstruction := value.KeySearch(batName).Value; batInstruction != nil {
				val := reflect.ValueOf(batInstruction).Interface().(*batType.Battery)
				value.SetValue(instructionData.FieldByName("Output").Interface().(string), GetBatteryState(val, finished))
			} else {
				fmt.Println("FIBER ERROR: Unable to find the battery named: ", batName)
				finished <- true
			}
		}()
		<-finished
	case "GetBatteryPercentage":
		go func() {
			batName := instructionData.FieldByName("BatteryName").Interface().(string)
			if batInstruction := value.KeySearch(batName).Value; batInstruction != nil {
				val := reflect.ValueOf(batInstruction).Interface().(*batType.Battery)
				value.SetValue(instructionData.FieldByName("Output").Interface().(string), GetBatteryPercentage(val, finished))
			} else {
				fmt.Println("FIBER ERROR: Unable to find the battery named: ", batName)
				finished <- true
			}
		}()
		<-finished
	case "GetBatteryRemainingTime":
		go func() {
			batName := instructionData.FieldByName("BatteryName").Interface().(string)
			if batInstruction := value.KeySearch(batName).Value; batInstruction != nil {
				val := reflect.ValueOf(batInstruction).Interface().(*batType.Battery)
				value.SetValue(instructionData.FieldByName("Output").Interface().(string), GetBatteryRemainingTime(val, finished))
			} else {
				fmt.Println("FIBER ERROR: Unable to find the battery named: ", batName)
				finished <- true
			}
		}()
		<-finished
	case "GetBatteryChargeRate":
		go func() {
			batName := instructionData.FieldByName("BatteryName").Interface().(string)
			if batInstruction := value.KeySearch(batName).Value; batInstruction != nil {
				val := reflect.ValueOf(batInstruction).Interface().(*batType.Battery)
				value.SetValue(instructionData.FieldByName("Output").Interface().(string), GetBatteryChargeRate(val, finished))
			} else {
				fmt.Println("FIBER ERROR: Unable to find the battery named: ", batName)
				finished <- true
			}
		}()
		<-finished
	case "GetBatteryCurrentCapacity":
		go func() {
			batName := instructionData.FieldByName("BatteryName").Interface().(string)
			if batInstruction := value.KeySearch(batName).Value; batInstruction != nil {
				val := reflect.ValueOf(batInstruction).Interface().(*batType.Battery)
				value.SetValue(instructionData.FieldByName("Output").Interface().(string), GetBatteryCurrentCapacity(val, finished))
			} else {
				fmt.Println("FIBER ERROR: Unable to find the battery named: ", batName)
				finished <- true
			}
		}()
		<-finished
	case "GetBatteryLastFullCapacity":
		go func() {
			batName := instructionData.FieldByName("BatteryName").Interface().(string)
			if batInstruction := value.KeySearch(batName).Value; batInstruction != nil {
				val := reflect.ValueOf(batInstruction).Interface().(*batType.Battery)
				value.SetValue(instructionData.FieldByName("Output").Interface().(string), GetBatteryLastFullCapacity(val, finished))
			} else {
				fmt.Println("FIBER ERROR: Unable to find the battery named: ", batName)
				finished <- true
			}
		}()
		<-finished
	case "GetBatteryDesignCapacity":
		go func() {
			batName := instructionData.FieldByName("BatteryName").Interface().(string)
			if batInstruction := value.KeySearch(batName).Value; batInstruction != nil {
				val := reflect.ValueOf(batInstruction).Interface().(*batType.Battery)
				value.SetValue(instructionData.FieldByName("Output").Interface().(string), GetBatteryDesignCapacity(val, finished))
			} else {
				fmt.Println("FIBER ERROR: Unable to find the battery named: ", batName)
				finished <- true
			}
		}()
		<-finished
	case "GetBatteryVoltage":
		go func() {
			batName := instructionData.FieldByName("BatteryName").Interface().(string)
			if batInstruction := value.KeySearch(batName).Value; batInstruction != nil {
				val := reflect.ValueOf(batInstruction).Interface().(*batType.Battery)
				value.SetValue(instructionData.FieldByName("Output").Interface().(string), GetBatteryVoltage(val, finished))
			} else {
				fmt.Println("FIBER ERROR: Unable to find the battery named: ", batName)
				finished <- true
			}
		}()
		<-finished
	case "GetBatteryDesignVoltage":
		go func() {
			batName := instructionData.FieldByName("BatteryName").Interface().(string)
			if batInstruction := value.KeySearch(batName).Value; batInstruction != nil {
				val := reflect.ValueOf(batInstruction).Interface().(*batType.Battery)
				value.SetValue(instructionData.FieldByName("Output").Interface().(string), GetBatteryDesignVoltage(val, finished))
			} else {
				fmt.Println("FIBER ERROR: Unable to find the battery named: ", batName)
				finished <- true
			}
		}()
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
}
