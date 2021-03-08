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
			newValue := &value.InstructionValue{
				Key:   instructionData.FieldByName("Output").Interface().(string),
				Value: GetBattery(finished),
			}
			value.FiberValues = append(value.FiberValues, newValue)
		}()
		<-finished
	case "GetBatteryState":
		go func() {
			batName := instructionData.FieldByName("BatteryName").Interface().(string)
			batInstruction := value.KeySearch(batName)
			if batInstruction != nil {
				val := reflect.ValueOf(batInstruction.Value).Interface().(*batType.Battery)
				newValue := &value.InstructionValue{
					Key:   instructionData.FieldByName("Output").Interface().(string),
					Value: GetBatteryState(val, finished),
				}
				value.FiberValues = append(value.FiberValues, newValue)
			} else {
				fmt.Println("FIBER: Unable to find the battery named: ", batName)
				finished <- true
			}
		}()
		<-finished
	case "GetBatteryPercentage":
		go func() {
			batName := instructionData.FieldByName("BatteryName").Interface().(string)
			batInstruction := value.KeySearch(batName)
			if batInstruction != nil {
				val := reflect.ValueOf(batInstruction.Value).Interface().(*batType.Battery)
				newValue := &value.InstructionValue{
					Key:   instructionData.FieldByName("Output").Interface().(string),
					Value: GetBatteryPercentage(val, finished),
				}
				value.FiberValues = append(value.FiberValues, newValue)
			} else {
				fmt.Println("FIBER: Unable to find the battery named: ", batName)
				finished <- true
			}
		}()
		<-finished
	case "GetBatteryRemainingTime":
		go func() {
			batName := instructionData.FieldByName("BatteryName").Interface().(string)
			batInstruction := value.KeySearch(batName)
			if batInstruction != nil {
				val := reflect.ValueOf(batInstruction.Value).Interface().(*batType.Battery)
				newValue := &value.InstructionValue{
					Key:   instructionData.FieldByName("Output").Interface().(string),
					Value: GetBatteryRemainingTime(val, finished),
				}
				value.FiberValues = append(value.FiberValues, newValue)
			} else {
				fmt.Println("FIBER: Unable to find the battery named: ", batName)
				finished <- true
			}
		}()
		<-finished
	case "GetBatteryChargeRate":
		go func() {
			batName := instructionData.FieldByName("BatteryName").Interface().(string)
			batInstruction := value.KeySearch(batName)
			if batInstruction != nil {
				val := reflect.ValueOf(batInstruction.Value).Interface().(*batType.Battery)
				newValue := &value.InstructionValue{
					Key:   instructionData.FieldByName("Output").Interface().(string),
					Value: GetBatteryChargeRate(val, finished),
				}
				value.FiberValues = append(value.FiberValues, newValue)
			} else {
				fmt.Println("FIBER: Unable to find the battery named: ", batName)
				finished <- true
			}
		}()
		<-finished
	case "GetBatteryCurrentCapacity":
		go func() {
			batName := instructionData.FieldByName("BatteryName").Interface().(string)
			batInstruction := value.KeySearch(batName)
			if batInstruction != nil {
				val := reflect.ValueOf(batInstruction.Value).Interface().(*batType.Battery)
				newValue := &value.InstructionValue{
					Key:   instructionData.FieldByName("Output").Interface().(string),
					Value: GetBatteryCurrentCapacity(val, finished),
				}
				value.FiberValues = append(value.FiberValues, newValue)
			} else {
				fmt.Println("FIBER: Unable to find the battery named: ", batName)
				finished <- true
			}
		}()
		<-finished
	case "GetBatteryLastFullCapacity":
		go func() {
			batName := instructionData.FieldByName("BatteryName").Interface().(string)
			batInstruction := value.KeySearch(batName)
			if batInstruction != nil {
				val := reflect.ValueOf(batInstruction.Value).Interface().(*batType.Battery)
				newValue := &value.InstructionValue{
					Key:   instructionData.FieldByName("Output").Interface().(string),
					Value: GetBatteryLastFullCapacity(val, finished),
				}
				value.FiberValues = append(value.FiberValues, newValue)
			} else {
				fmt.Println("FIBER: Unable to find the battery named: ", batName)
				finished <- true
			}
		}()
		<-finished
	case "GetBatteryDesignCapacity":
		go func() {
			batName := instructionData.FieldByName("BatteryName").Interface().(string)
			batInstruction := value.KeySearch(batName)
			if batInstruction != nil {
				val := reflect.ValueOf(batInstruction.Value).Interface().(*batType.Battery)
				newValue := &value.InstructionValue{
					Key:   instructionData.FieldByName("Output").Interface().(string),
					Value: GetBatteryDesignCapacity(val, finished),
				}
				value.FiberValues = append(value.FiberValues, newValue)
			} else {
				fmt.Println("FIBER: Unable to find the battery named: ", batName)
				finished <- true
			}
		}()
		<-finished
	case "GetBatteryVoltage":
		go func() {
			batName := instructionData.FieldByName("BatteryName").Interface().(string)
			batInstruction := value.KeySearch(batName)
			if batInstruction != nil {
				val := reflect.ValueOf(batInstruction.Value).Interface().(*batType.Battery)
				newValue := &value.InstructionValue{
					Key:   instructionData.FieldByName("Output").Interface().(string),
					Value: GetBatteryVoltage(val, finished),
				}
				value.FiberValues = append(value.FiberValues, newValue)
			} else {
				fmt.Println("FIBER: Unable to find the battery named: ", batName)
				finished <- true
			}
		}()
		<-finished
	case "GetBatteryDesignVoltage":
		go func() {
			batName := instructionData.FieldByName("BatteryName").Interface().(string)
			batInstruction := value.KeySearch(batName)
			if batInstruction != nil {
				val := reflect.ValueOf(batInstruction.Value).Interface().(*batType.Battery)
				newValue := &value.InstructionValue{
					Key:   instructionData.FieldByName("Output").Interface().(string),
					Value: GetBatteryDesignVoltage(val, finished),
				}
				value.FiberValues = append(value.FiberValues, newValue)
			} else {
				fmt.Println("FIBER: Unable to find the battery named: ", batName)
				finished <- true
			}
		}()
		<-finished
	default:
		fmt.Println("FIBER: This function is not integrated yet: " + funcName)
	}
}
