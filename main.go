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
			flights := io.GetRange(0, 7)

			for i, fl := range flights {
				ui.TblBox.SetCell(i+2, 0, fl.TimeFrom)
				ui.TblBox.SetCell(i+2, 1, fl.FlightFrom)
				ui.TblBox.SetCell(i+2, 2, fl.FlightTo)
				ui.TblBox.SetCell(i+2, 3, fl.TimeTo)
			}

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
