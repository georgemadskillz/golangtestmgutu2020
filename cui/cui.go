package cui

import (
	"fmt"
	"log"
	"strconv"

	"golang.org/x/crypto/ssh/terminal"
)

// UIctrl is a common structure handling UI
type UIctrl struct {
	Scr       Screen
	TermState *terminal.State
	Extbox    Box
	CommonBox InfoBox
	StatusBox InfoBox
	TblBox    TableBox
}

// Init initialize UI controller
func (ui *UIctrl) Init() {
	ui.Scr.Init()

	state, err := terminal.MakeRaw(0)
	if err != nil {
		//log.Fatalln("setting stdin to raw:", err)
	}
	ui.TermState = state

	// Ext border
	ui.Extbox = Box{0, 0, ui.Scr.Width, ui.Scr.Height, false}

	// Common window
	ui.CommonBox = InfoBox{1, 1, 40, 10, false, nil}
	ui.CommonBox.Init()
	ui.CommonBox.SetLineText(0, "БАЗА ДАННЫХ")
	ui.CommonBox.SetLineText(2, "Выбор таблицы:")
	ui.CommonBox.SetLineText(4, "# Рейсы")
	ui.CommonBox.SetLineText(5, "# Аэропорты")
	ui.CommonBox.SetLineText(6, "# Цены")
	ui.CommonBox.SetActiveState(true)

	// Table window
	ui.TblBox = TableBox{41, 1, ui.Scr.Width - 2 - 40, ui.Scr.Height - 2, false, 4, nil, nil}
	ui.TblBox.Init()
	// ui.TblBox.SetColsWidths(10, 10, 10, 15)
	// ui.TblBox.SetCell(0, 0, "Date from")
	// ui.TblBox.SetCell(0, 1, "From")
	// ui.TblBox.SetCell(0, 2, "To")
	// ui.TblBox.SetCell(0, 3, "Date to")

	log.Printf("Cells[0] before=[%#v]\r", ui.TblBox.Cells[0])
	str := ""

	str = "my w=" + strconv.Itoa(ui.TblBox.CellsWidths[0])
	ui.TblBox.SetCell(0, 0, str)
	str = "my w=" + strconv.Itoa(ui.TblBox.CellsWidths[0])
	ui.TblBox.SetCell(0, 1, str)
	str = "my w=" + strconv.Itoa(ui.TblBox.CellsWidths[0])
	ui.TblBox.SetCell(0, 2, str)
	str = "my w=" + strconv.Itoa(ui.TblBox.CellsWidths[0])
	ui.TblBox.SetCell(0, 3, str)

	// Status window
	ui.StatusBox = InfoBox{1, 11, 40, ui.Scr.Height - 2 - 10, false, nil}
	ui.StatusBox.Init()
	ui.StatusBox.SetLineText(0, "Статус программы:")

	ui.StatusBox.SetLineText(2, "Текущая таблица: <Рейсы>")

	ui.StatusBox.SetLineText(14, "Управление программой:")
	ui.StatusBox.SetLineText(15, "Esc: выход из программы")
	ui.StatusBox.SetLineText(16, "Tab: переход между окнами")

	ui.StatusBox.SetLineText(18, "─────────────────────────────────────")
	ui.StatusBox.SetLineText(19, "Отладочная информация:")

	ui.Scr.UpdateSize()
	ui.Draw(&ui.Scr)
	ui.Scr.SendToDisplay()
}

// Draw draws all it's elements
func (ui *UIctrl) Draw(scr *Screen) {
	ui.Extbox.Draw(scr)
	ui.CommonBox.Draw(scr)
	ui.StatusBox.Draw(scr)
	ui.TblBox.Draw(scr)
}

// DeInit restores terminal state from raw and clears screen
func (ui *UIctrl) DeInit() {
	if err := terminal.Restore(0, ui.TermState); err != nil {
		//log.Println("warning, failed to restore terminal:", err)
	}
	ui.Scr.Clear()
}

var debugLine int = 20

// DebugPrintln is
func (ui *UIctrl) DebugPrintln(str string, a ...interface{}) {
	msg := fmt.Sprintf(str, a)
	ui.StatusBox.SetLineText(debugLine, msg)
	debugLine++
	if debugLine >= len(ui.StatusBox.Text) {
		debugLine = 20
	}
}
