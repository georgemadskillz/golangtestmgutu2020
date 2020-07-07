package appctrl

import (
	"flydb/cui"
	"flydb/datamdl"
	"flydb/ioctrl"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// AppController is application main controller
type AppController struct {
	termState *terminal.State
	io        ioctrl.FlyDbIO
	ui        cui.UIctrl
}

// Init is
func (app *AppController) Init() {
	app.termState, _ = terminal.MakeRaw(0)

	app.ui.Init()
	app.io.Init()
}

// Exit is
func (app *AppController) Exit() {
	app.ui.DeInit()

	terminal.Restore(0, app.termState)
	os.Exit(0)
}

// CheckKbdInput is
func (app *AppController) CheckKbdInput(kbdKey rune) {
	switch kbdKey {
	case '\x1b': // Esc
		app.Exit()
	case '\t': // Tab

		if app.ui.CommonBox.IsActive {
			app.ui.CommonBox.SetActiveState(false)
			app.ui.TblBox.SetActiveState(true)
		} else {
			app.ui.CommonBox.SetActiveState(true)
			app.ui.TblBox.SetActiveState(false)
		}

		app.ui.Draw(&app.ui.Scr)
		app.ui.Scr.SendToDisplay()
	case '1':
		flights := app.io.GetRange(ioctrl.FdbFly, 0, 5)

		for i, fl := range flights {
			f := fl.(datamdl.Flight)
			app.ui.TblBox.SetCell(i+2, 0, f.TimeFrom)
			app.ui.TblBox.SetCell(i+2, 1, f.FlightFrom)
			app.ui.TblBox.SetCell(i+2, 2, f.FlightTo)
			app.ui.TblBox.SetCell(i+2, 3, f.TimeTo)
		}
		app.ui.Draw(&app.ui.Scr)
		app.ui.Scr.SendToDisplay()
	case '2':
		airports := app.io.GetRange(ioctrl.FdbAir, 0, 5)

		for i, air := range airports {
			a := air.(datamdl.Airport)
			app.ui.TblBox.SetCell(i+2, 0, a.AirID)
			app.ui.TblBox.SetCell(i+2, 1, a.AirCity)
			app.ui.TblBox.SetCell(i+2, 2, a.AirName)
		}
		app.ui.Draw(&app.ui.Scr)
		app.ui.Scr.SendToDisplay()
	case '3':
		prices := app.io.GetRange(ioctrl.FdbPrc, 0, 5)

		for i, pr := range prices {
			p := pr.(datamdl.Price)
			app.ui.TblBox.SetCell(i+2, 0, p.FlightID)
			app.ui.TblBox.SetCell(i+2, 1, p.Seat)
			app.ui.TblBox.SetCell(i+2, 2, p.Price)
		}
		app.ui.Draw(&app.ui.Scr)
		app.ui.Scr.SendToDisplay()
	default:
		app.ui.Scr.SendToDisplay()
	}
}

// // DebugPrintln is
// func (ui *UIctrl) DebugPrintln(str string, a ...interface{}) {
// 	msg := fmt.Sprintf(str, a)
// 	ui.StatusBox.SetLineText(debugLine, msg)
// 	debugLine++
// 	if debugLine >= len(ui.StatusBox.Text) {
// 		debugLine = 20
// 	}
// }
