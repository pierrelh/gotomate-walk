package battery

import (
	"fmt"
	"os"
	"time"
)

// GetAll returns information about all batteries in the system.
// If error != nil, it will be either ErrFatal or Errors.
// If error is of type Errors, it is guaranteed that length of both returned slices is the same and that i-th error coresponds with i-th battery structure.
func GetAll() ([]*Battery, error) {
	return getAll(systemGetAll)
}

// GetBattery get the first system battery if exist
func GetBattery(finished chan bool) *Battery {
	fmt.Println("Battery getting initialization ...")
	batteries, err := GetAll()
	if _, isFatal := err.(ErrFatal); isFatal {
		finished <- true
		return nil
	}
	if len(batteries) == 0 {
		finished <- true
		return nil
	}
	errs, partialErrs := err.(Errors)
	for i, bat := range batteries {
		if partialErrs && errs[i] != nil {
			fmt.Fprintf(os.Stderr, "Error getting info for BAT%d: %s\n", i, errs[i])
			continue
		}
		finished <- true
		return bat
	}
	finished <- true
	return nil
}

// GetBatteryState : Return the battery state of a battery
func GetBatteryState(bat *Battery) State {
	return bat.State
}

// GetBatteryPercentage : Return the left percentage of a battery
func GetBatteryPercentage(bat *Battery) float64 {
	return bat.Current / bat.Full * 100
}

// GetBatteryRemainingTime : Return the remaining time of battery or for battery charging
func GetBatteryRemainingTime(bat *Battery) time.Duration {
	var timeNum float64
	switch bat.State {
	case Discharging:
		timeNum = bat.Current / bat.ChargeRate
	case Charging:
		timeNum = (bat.Full - bat.Current) / bat.ChargeRate
	default:
		timeNum = 0
	}
	duration, _ := time.ParseDuration(fmt.Sprintf("%fh", timeNum))
	return duration
}

// GetBatteryChargeRate : Return a battery charge rate mW
func GetBatteryChargeRate(bat *Battery) time.Duration {
	return time.Duration(bat.ChargeRate)
}

// GetBatteryCurrentCapacity : Return a battery current capacity
func GetBatteryCurrentCapacity(bat *Battery) float64 {
	return bat.Current
}

// GetBatteryLastFullCapacity : Return a battery last full capacity
func GetBatteryLastFullCapacity(bat *Battery) float64 {
	return bat.Full
}

// GetBatteryDesignCapacity : Return a battery design capacity
func GetBatteryDesignCapacity(bat *Battery) float64 {
	return bat.Design
}

// GetBatteryVoltage : Return a battery voltage
func GetBatteryVoltage(bat *Battery) float64 {
	return bat.Voltage
}

// GetBatteryDesignVoltage : Return a battery design voltage
func GetBatteryDesignVoltage(bat *Battery) float64 {
	return bat.DesignVoltage
}
