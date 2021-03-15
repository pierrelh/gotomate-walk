package battery

import (
	"fmt"
	"os"
	"time"

	"github.com/distatus/battery"
)

// GetBattery get the first system battery if exist
func GetBattery(finished chan bool) *battery.Battery {
	fmt.Println("FIBER: Getting battery ...")
	batteries, err := battery.GetAll()
	if _, isFatal := err.(battery.ErrFatal); isFatal {
		finished <- true
		return nil
	}
	if len(batteries) == 0 {
		finished <- true
		return nil
	}
	errs, partialErrs := err.(battery.Errors)
	for i, bat := range batteries {
		if partialErrs && errs[i] != nil {
			fmt.Fprintf(os.Stderr, "FIBER: Error getting info for BAT%d: %s\n", i, errs[i])
			continue
		}
		finished <- true
		return bat
	}
	finished <- true
	return nil
}

// GetBatteryState : Return the battery state of a battery
func GetBatteryState(bat *battery.Battery, finished chan bool) battery.State {
	fmt.Println("FIBER: Getting battery state ...")
	finished <- true
	return bat.State
}

// GetBatteryPercentage : Return the left percentage of a battery
func GetBatteryPercentage(bat *battery.Battery, finished chan bool) float64 {
	fmt.Println("FIBER: Getting battery percentage ...")
	finished <- true
	return bat.Current / bat.Full * 100
}

// GetBatteryRemainingTime : Return the remaining time of battery or for battery charging
func GetBatteryRemainingTime(bat *battery.Battery, finished chan bool) time.Duration {
	fmt.Println("FIBER: Getting battery remaining time ...")
	var timeNum float64
	switch bat.State {
	case battery.Discharging:
		timeNum = bat.Current / bat.ChargeRate
	case battery.Charging:
		timeNum = (bat.Full - bat.Current) / bat.ChargeRate
	default:
		timeNum = 0
	}
	duration, _ := time.ParseDuration(fmt.Sprintf("%fh", timeNum))
	finished <- true
	return duration
}

// GetBatteryChargeRate : Return a battery charge rate mW
func GetBatteryChargeRate(bat *battery.Battery, finished chan bool) time.Duration {
	fmt.Println("FIBER: Getting battery charge rate ...")
	finished <- true
	return time.Duration(bat.ChargeRate)
}

// GetBatteryCurrentCapacity : Return a battery current capacity
func GetBatteryCurrentCapacity(bat *battery.Battery, finished chan bool) float64 {
	fmt.Println("FIBER: Getting battery current capacity ...")
	finished <- true
	return bat.Current
}

// GetBatteryLastFullCapacity : Return a battery last full capacity
func GetBatteryLastFullCapacity(bat *battery.Battery, finished chan bool) float64 {
	fmt.Println("FIBER: Getting battery last full capacity ...")
	finished <- true
	return bat.Full
}

// GetBatteryDesignCapacity : Return a battery design capacity
func GetBatteryDesignCapacity(bat *battery.Battery, finished chan bool) float64 {
	fmt.Println("FIBER: Getting battery design capacity ...")
	finished <- true
	return bat.Design
}

// GetBatteryVoltage : Return a battery voltage
func GetBatteryVoltage(bat *battery.Battery, finished chan bool) float64 {
	fmt.Println("FIBER: Getting battery voltage ...")
	finished <- true
	return bat.Voltage
}

// GetBatteryDesignVoltage : Return a battery design voltage
func GetBatteryDesignVoltage(bat *battery.Battery, finished chan bool) float64 {
	fmt.Println("FIBER: Getting battery design voltage ...")
	finished <- true
	return bat.DesignVoltage
}