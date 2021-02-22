package automate

import (
	"fmt"
	"gotomate/convert"
	"gotomate/systime"
)

// LaunchAutomate Start the fiber
func LaunchAutomate() {
	sysTime := systime.GetCurrentSysTime()
	fmt.Println(sysTime)

	h, m, s := systime.GetCurrentSysClock()
	var a []int
	a = append(a, h, m, s)
	stringSysClock := convert.IntArrayToString(a, " ")
	fmt.Println(stringSysClock)

	// notifActions := []toast.Action{}

	// newNotification := notification.Create("Test", "Test MSG", notifActions)

	// err := notification.Push(newNotification)
	// if err != nil {
	// 	fmt.Println("An error occur when pushing notification")
	// }

	// bat := battery.GetBattery()
	// if bat != nil {
	// 	fmt.Println(battery.GetBatteryPercentage(bat))
	// 	fmt.Println(battery.GetBatteryRemainingTime(bat))
	// 	fmt.Println(battery.GetBatteryChargeRate(bat))
	// }

	// clipboard.Write("TestClipboard")

	// clipboardValue, err := clipboard.Read()
	// log.Print(clipboardValue)

	// mouse.Move(250, 250)
	// keyboard.Tap("a")

}
