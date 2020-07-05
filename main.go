package main

import (
	"bufio"
	"flydb/cui"
	"flydb/datamdl"
	"flydb/ioctrl"
	"log"
	"os"
)

func main() {

	ui := cui.UIctrl{}
	ui.Init()

	io := ioctrl.FlyDbIO{}
	io.CuiPtr = &ui
	io.Init()

	inp := bufio.NewReader(os.Stdin)
	for {

		switch r, _, _ := inp.ReadRune(); r {
		case '\x1b': // Esc
			ui.DeInit()
			os.Exit(0)
		case '\t': // Tab

			if ui.CommonBox.IsActive {
				ui.CommonBox.SetActiveState(false)
				ui.TblBox.SetActiveState(true)
			} else {
				ui.CommonBox.SetActiveState(true)
				ui.TblBox.SetActiveState(false)
			}

			ui.Draw(&ui.Scr)
			ui.Scr.SendToDisplay()
		case '1':
			flights := io.GetRange(ioctrl.FdbFly, 0, 5)

			log.Printf("Got flights=[%#v]\r", flights)

			for i, fl := range flights {
				f := fl.(datamdl.Flight)
				ui.TblBox.SetCell(i+2, 0, f.TimeFrom)
				ui.TblBox.SetCell(i+2, 1, f.FlightFrom)
				ui.TblBox.SetCell(i+2, 2, f.FlightTo)
				ui.TblBox.SetCell(i+2, 3, f.TimeTo)
			}
		case '2':
		case '3':
		default:
			ui.Scr.SendToDisplay()
		}
	}
}
