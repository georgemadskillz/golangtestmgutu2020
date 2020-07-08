package appctrl

import (
	"flydb/cui"
	"flydb/datamdl"
	"flydb/ioctrl"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// AppController is application main controller
type AppController struct {
	termState *terminal.State
	scr       cui.Screen
	io        ioctrl.FlyDbIO

	Extbox    cui.Box
	CommonBox cui.InfoBox
	StatusBox cui.InfoBox

	TblBox   cui.TableBox
	currTbl  int
	tblNames []string

	debugLine int
}

// Init is
func (app *AppController) Init() {
	app.termState, _ = terminal.MakeRaw(0)
	app.scr.Init()
	app.io.Init()

	app.debugLine = 20

	app.tblNames = make([]string, 3, 3)
	app.tblNames[ioctrl.FdbFly] = "<Полеты>"
	app.tblNames[ioctrl.FdbAir] = "<Аэропорты>"
	app.tblNames[ioctrl.FdbPrc] = "<Цены>"
	app.SetCurrTbl(ioctrl.FdbFly)

	// Ext border
	app.Extbox = cui.Box{0, 0, app.scr.Width, app.scr.Height, false}

	// Common window
	app.CommonBox = cui.InfoBox{1, 1, 40, 10, false, nil}
	app.CommonBox.Init()
	app.CommonBox.SetLineText(0, "БАЗА ДАННЫХ")
	app.CommonBox.SetLineText(2, "Выбор таблицы:")
	app.CommonBox.SetLineText(4, "# Рейсы")
	app.CommonBox.SetLineText(5, "# Аэропорты")
	app.CommonBox.SetLineText(6, "# Цены")
	app.CommonBox.SetActiveState(true)

	// Status window
	app.StatusBox = cui.InfoBox{1, 11, 40, app.scr.Height - 2 - 10, false, nil}
	app.StatusBox.Init()
	app.StatusBox.SetLineText(0, "Статус программы:")
	app.StatusBox.SetLineText(1, "─────────────────────────────────────")
	app.StatusBox.SetLineText(2, "Текущая таблица: "+app.tblNames[app.currTbl])
	app.StatusBox.SetLineText(14, "Управление программой:")
	app.StatusBox.SetLineText(15, "Esc: выход из программы")
	app.StatusBox.SetLineText(16, "Tab: переход между окнами")

	app.StatusBox.SetLineText(18, "─────────────────────────────────────")
	app.StatusBox.SetLineText(19, "Отладочная информация:")

	// Table window
	app.TblBox = cui.TableBox{41, 1, app.scr.Width - 2 - 40, app.scr.Height - 2, false, 4, nil, nil}
	app.TblBox.Init()
	app.TblBox.SetCell(0, 0, "Date from")
	app.TblBox.SetCell(0, 1, "From")
	app.TblBox.SetCell(0, 2, "To")
	app.TblBox.SetCell(0, 3, "Date to")

	app.TblBox.FillCell(1, 0, '═')
	app.TblBox.FillCell(1, 1, '═')
	app.TblBox.FillCell(1, 2, '═')
	app.TblBox.FillCell(1, 3, '═')

	app.Draw()
}

// Exit is
func (app *AppController) Exit() {
	app.scr.Clear()
	terminal.Restore(0, app.termState)
	os.Exit(0)
}

// SetCurrTbl is
func (app *AppController) SetCurrTbl(tblKey int) {
	if tblKey >= ioctrl.FdbAmount {
		fmt.Errorf("unknown table key, set default table ioctrl.FdbFly")
	}

	app.currTbl = tblKey
	app.StatusBox.SetLineText(2, "Текущая таблица: "+app.tblNames[tblKey])
}

// Draw is
func (app *AppController) Draw() {
	app.Extbox.Draw(&app.scr)
	app.CommonBox.Draw(&app.scr)
	app.StatusBox.Draw(&app.scr)
	app.TblBox.Draw(&app.scr)

	app.scr.SendToDisplay()
}

// CheckKbdInput is
func (app *AppController) CheckKbdInput(kbdKey rune) {
	switch kbdKey {
	case '\x1b': // Esc
		app.Exit()
	case '\t': // Tab

		if app.CommonBox.IsActive {
			app.CommonBox.SetActiveState(false)
			app.TblBox.SetActiveState(true)
		} else {
			app.CommonBox.SetActiveState(true)
			app.TblBox.SetActiveState(false)
		}

		app.Draw()
	case '1':
		app.SetCurrTbl(ioctrl.FdbFly)

		flights := app.io.GetRange(ioctrl.FdbFly, 0, 5)

		for i, fl := range flights {
			f := fl.(datamdl.Flight)
			app.TblBox.SetCell(i+2, 0, f.TimeFrom)
			app.TblBox.SetCell(i+2, 1, f.FlightFrom)
			app.TblBox.SetCell(i+2, 2, f.FlightTo)
			app.TblBox.SetCell(i+2, 3, f.TimeTo)
		}
		app.Draw()
	case '2':
		app.SetCurrTbl(ioctrl.FdbAir)

		airports := app.io.GetRange(ioctrl.FdbAir, 0, 5)

		for i, air := range airports {
			a := air.(datamdl.Airport)
			app.TblBox.SetCell(i+2, 0, a.AirID)
			app.TblBox.SetCell(i+2, 1, a.AirCity)
			app.TblBox.SetCell(i+2, 2, a.AirName)
		}
		app.Draw()
	case '3':
		app.SetCurrTbl(ioctrl.FdbPrc)

		prices := app.io.GetRange(ioctrl.FdbPrc, 0, 5)

		for i, pr := range prices {
			p := pr.(datamdl.Price)
			app.TblBox.SetCell(i+2, 0, p.FlightID)
			app.TblBox.SetCell(i+2, 1, p.Seat)
			app.TblBox.SetCell(i+2, 2, p.Price)
		}
		app.Draw()
	default:
		app.Draw()
	}
}

// DebugPrintln is
func (app *AppController) DebugPrintln(str string, a ...interface{}) {
	msg := fmt.Sprintf("%v"+str, a)
	app.StatusBox.SetLineText(app.debugLine, msg)
	app.debugLine++
	if app.debugLine >= len(app.StatusBox.Text) {
		app.debugLine = 20
	}
}
