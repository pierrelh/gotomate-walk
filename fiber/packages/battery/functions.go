package battery

import (
	"fmt"
	"gotomate/fiber/variable"
	"os"
	"reflect"
	"time"

	"github.com/distatus/battery"
	batType "github.com/distatus/battery"
)

// GetBattery get the first system battery if exist
func GetBattery(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery ...")
	batteries, err := battery.GetAll()
	if _, isFatal := err.(battery.ErrFatal); !isFatal {
		if len(batteries) != 0 {
			errs, partialErrs := err.(battery.Errors)
			for i, bat := range batteries {
				if partialErrs && errs[i] != nil {
					fmt.Fprintf(os.Stderr, "FIBER ERROR: Error getting info for BAT%d: %s\n", i, errs[i])
					continue
				}
				variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), bat)
				finished <- true
				return -1
			}
		}
	}
	finished <- true
	return -1
}

// GetBatteryChargeRate : Return a battery charge rate mW
func GetBatteryChargeRate(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery charge rate ...")
	batName := instructionData.FieldByName("BatteryName").Interface().(string)
	if batInstruction := variable.SearchVariable(batName).Value; batInstruction != nil {
		bat := reflect.ValueOf(batInstruction).Interface().(*batType.Battery)
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), time.Duration(bat.ChargeRate))
	} else {
		fmt.Println("FIBER ERROR: Unable to find the battery named: ", batName)
	}
	finished <- true
	return -1
}

// GetBatteryCurrentCapacity : Return a battery current capacity
func GetBatteryCurrentCapacity(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery current capacity ...")
	batName := instructionData.FieldByName("BatteryName").Interface().(string)
	if batInstruction := variable.SearchVariable(batName).Value; batInstruction != nil {
		bat := reflect.ValueOf(batInstruction).Interface().(*batType.Battery)
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), bat.Current)
	} else {
		fmt.Println("FIBER ERROR: Unable to find the battery named: ", batName)
	}
	finished <- true
	return -1
}

// GetBatteryDesignCapacity : Return a battery design capacity
func GetBatteryDesignCapacity(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery design capacity ...")
	batName := instructionData.FieldByName("BatteryName").Interface().(string)
	if batInstruction := variable.SearchVariable(batName).Value; batInstruction != nil {
		bat := reflect.ValueOf(batInstruction).Interface().(*batType.Battery)
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), bat.Design)
	} else {
		fmt.Println("FIBER ERROR: Unable to find the battery named: ", batName)
	}
	finished <- true
	return -1
}

// GetBatteryDesignVoltage : Return a battery design voltage
func GetBatteryDesignVoltage(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery design voltage ...")
	batName := instructionData.FieldByName("BatteryName").Interface().(string)
	if batInstruction := variable.SearchVariable(batName).Value; batInstruction != nil {
		bat := reflect.ValueOf(batInstruction).Interface().(*batType.Battery)
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), bat.DesignVoltage)
	} else {
		fmt.Println("FIBER ERROR: Unable to find the battery named: ", batName)
	}
	finished <- true
	return -1
}

// GetBatteryLastFullCapacity : Return a battery last full capacity
func GetBatteryLastFullCapacity(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery last full capacity ...")
	batName := instructionData.FieldByName("BatteryName").Interface().(string)
	if batInstruction := variable.SearchVariable(batName).Value; batInstruction != nil {
		bat := reflect.ValueOf(batInstruction).Interface().(*batType.Battery)
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), bat.Full)
	} else {
		fmt.Println("FIBER ERROR: Unable to find the battery named: ", batName)
	}
	finished <- true
	return -1
}

// GetBatteryPercentage : Return the left percentage of a battery
func GetBatteryPercentage(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery percentage ...")
	batName := instructionData.FieldByName("BatteryName").Interface().(string)
	if batInstruction := variable.SearchVariable(batName).Value; batInstruction != nil {
		bat := reflect.ValueOf(batInstruction).Interface().(*batType.Battery)
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), bat.Current/bat.Full*100)
	} else {
		fmt.Println("FIBER ERROR: Unable to find the battery named: ", batName)
	}
	finished <- true
	return -1
}

// GetBatteryRemainingTime : Return the remaining time of battery or for battery charging
func GetBatteryRemainingTime(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery remaining time ...")
	var timeNum float64
	batName := instructionData.FieldByName("BatteryName").Interface().(string)
	if batInstruction := variable.SearchVariable(batName).Value; batInstruction != nil {
		bat := reflect.ValueOf(batInstruction).Interface().(*batType.Battery)
		switch bat.State {
		case battery.Discharging:
			timeNum = bat.Current / bat.ChargeRate
		case battery.Charging:
			timeNum = (bat.Full - bat.Current) / bat.ChargeRate
		default:
			timeNum = 0
		}
		duration, _ := time.ParseDuration(fmt.Sprintf("%fh", timeNum))
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), duration)
	} else {
		fmt.Println("FIBER ERROR: Unable to find the battery named: ", batName)
	}
	finished <- true
	return -1
}

// GetBatteryState : Return the battery state of a battery
func GetBatteryState(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery state ...")
	batName := instructionData.FieldByName("BatteryName").Interface().(string)
	if batInstruction := variable.SearchVariable(batName).Value; batInstruction != nil {
		bat := reflect.ValueOf(batInstruction).Interface().(*batType.Battery)
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), bat.State)
	} else {
		fmt.Println("FIBER ERROR: Unable to find the battery named: ", batName)
	}
	finished <- true
	return -1
}

// GetBatteryVoltage : Return a battery voltage
func GetBatteryVoltage(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery voltage ...")
	batName := instructionData.FieldByName("BatteryName").Interface().(string)
	if batInstruction := variable.SearchVariable(batName).Value; batInstruction != nil {
		bat := reflect.ValueOf(batInstruction).Interface().(*batType.Battery)
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), bat.Voltage)
	} else {
		fmt.Println("FIBER ERROR: Unable to find the battery named: ", batName)
	}
	finished <- true
	return -1
}
