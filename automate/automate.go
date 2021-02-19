package automate

import (
	"gotomate/battery"
	"gotomate/convert"
	"gotomate/log"
	"gotomate/notification"
	"gotomate/systime"

	toast "gopkg.in/toast.v1"
)

// LaunchAutomate Start the fiber
func LaunchAutomate() {
	sysTime := systime.GetCurrentSysTime()
	log.Print(sysTime)

	h, m, s := systime.GetCurrentSysClock()
	var a []int
	a = append(a, h, m, s)
	stringSysClock := convert.IntArrayToString(a, " ")
	log.Print(stringSysClock)

	notifActions := []toast.Action{}

	newNotification := notification.Create("Test", "Test MSG", notifActions, false)

	err := notification.Push(newNotification)
	if err != nil {
		log.Print("An error occur when pushing notification")
	}

	bat := battery.GetBatteries()
	if bat != nil {
		log.Print(battery.GetBatteryPercentage(bat))
		log.Print(battery.GetBatteryRemainingTime(bat))
		log.Print(battery.GetBatteryChargeRate(bat))
	}

	// clipboard.Write("TestClipboard")

	// clipboardValue, err := clipboard.Read()
	// log.Print(clipboardValue)

	// mouse.Move(250, 250)
	// keyboard.Tap("a")

}
