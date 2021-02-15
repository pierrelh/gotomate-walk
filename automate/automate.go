package automate

import (
	"gotomate/battery"
	"gotomate/clipboard"
	"gotomate/convert"
	"gotomate/log"
	"gotomate/notification"
	"gotomate/systime"

	"github.com/go-vgo/robotgo"

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

	newNotification := notification.CreateNotification("Test", "Test MSG", notifActions, false)

	err := notification.PushNotification(newNotification)
	if err != nil {
		log.Print("An error occur when pushing notification")
	}

	bat := battery.GetBatteries()
	if bat != nil {
		log.Print(battery.GetBatteryPercentage(bat))
		log.Print(battery.GetBatteryRemainingTime(bat))
		log.Print(battery.GetBatteryChargeRate(bat))
	}

	clipboard.WriteClipboard("TestClipboard")

	clipboardValue, err := clipboard.ReadClipboard()
	log.Print(clipboardValue)

	robotgo.MoveMouse(250, 250)
	robotgo.KeyTap("a")

}
