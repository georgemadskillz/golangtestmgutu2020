package main

import (
	"bufio"
	"flydb/cui"
	"flydb/ioctrl"
	"os"
)

func main() {

	ui := cui.UIctrl{}
	ui.Init()

	io := ioctrl.FlyDbIO{}
	io.CuiPtr = &ui
	io.Init("database/flights.fdb", "database/airports.fdb", "database/prices.fdb")
	io.LoadFlyTable()

	inp := bufio.NewReader(os.Stdin)
	for {

		switch r, _, _ := inp.ReadRune(); r {
		case '\x1b': // Esc
			ui.DeInit()
			os.Exit(0)
		case '\t': // Tab
			// flights := io.GetRange(0, 7)

			// for i, fl := range flights {
			// 	str := fmt.Sprintf("%v|%v|%v|%v", fl.TimeFrom, fl.FlightFrom, fl.FlightTo, fl.TimeTo)
			// 	ui.TableBox.SetLineText(i, str)
			// }

			if ui.CommonBox.IsActive {
				ui.CommonBox.SetActiveState(false)
				ui.TblBox.SetActiveState(true)
			} else {
				ui.CommonBox.SetActiveState(true)
				ui.TblBox.SetActiveState(false)
			}

			ui.Draw(&ui.Scr)
			ui.Scr.SendToDisplay()
		default:
			ui.Scr.SendToDisplay()
		}
	}
}
