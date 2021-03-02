package battery

import (
	"time"

	"github.com/distatus/battery"
)

// UserBatteryDatabinder Define the UserBattery parameters
type UserBatteryDatabinder struct {
	Var   string
	Value *battery.Battery
}

// StateDatabinder Define the BatteryState parameters
type StateDatabinder struct {
	BatteryName string
	Var         string
	Value       battery.State
}

// PercentageDatabinder Define the BatteryPercentage parameters
type PercentageDatabinder struct {
	BatteryName string
	Var         string
	Value       float64
}

// RemainingTimeDatabinder Define the BatteryRemainingTime parameters
type RemainingTimeDatabinder struct {
	BatteryName string
	Var         string
	Value       time.Duration
}

// ChargeRateDatabinder Define the BatteryChargeRate parameters
type ChargeRateDatabinder struct {
	BatteryName string
	Var         string
	Value       time.Duration
}

// CurrentCapacityDatabinder Define the BatteryCurrentCapacity parameters
type CurrentCapacityDatabinder struct {
	BatteryName string
	Var         string
	Value       float64
}

// LastFullCapacityDatabinder Define the BatteryLastFullCapacity parameters
type LastFullCapacityDatabinder struct {
	BatteryName string
	Var         string
	Value       float64
}

// DesignCapacityDatabinder Define the BatteryDesignCapacity parameters
type DesignCapacityDatabinder struct {
	BatteryName string
	Var         string
	Value       float64
}

// VoltageDatabinder Define the BatteryVoltage parameters
type VoltageDatabinder struct {
	BatteryName string
	Var         string
	Value       float64
}

// DesignVoltageDatabinder Define the BatteryDesignVoltage parameters
type DesignVoltageDatabinder struct {
	BatteryName string
	Var         string
	Value       float64
}
