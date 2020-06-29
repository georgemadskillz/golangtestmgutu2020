package cui

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// UIctrl is a common structure handling UI
type UIctrl struct {
	Scr       Screen
	TermState *terminal.State
	Extbox    Box
	CommonBox InfoBox
	StatusBox InfoBox
	TableBox  InfoBox
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
	ui.TableBox = InfoBox{41, 1, ui.Scr.Width - 2 - 40, ui.Scr.Height - 2, false, nil}
	ui.TableBox.Init()
	ui.TableBox.SetLineText(0, "Здесь будет таблица базы данных")

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

	ui.Draw(&ui.Scr)
}

// Draw draws all it's elements
func (ui *UIctrl) Draw(scr *Screen) {
	ui.Extbox.Draw(scr)
	ui.CommonBox.Draw(scr)
	ui.StatusBox.Draw(scr)
	ui.TableBox.Draw(scr)
}

// DeInit restores terminal state from raw and clears screen
func (ui *UIctrl) DeInit() {
	if err := terminal.Restore(0, ui.TermState); err != nil {
		//log.Println("warning, failed to restore terminal:", err)
	}
	ui.Scr.Clear()
}

var debugLine int = 20

func (ui *UIctrl) debugPrintf(str string, a ...interface{}) {
	msg := fmt.Sprintf(str, a)
	ui.StatusBox.SetLineText(debugLine, msg)
	debugLine++
	if debugLine >= len(ui.StatusBox.Text) {
		debugLine = 20
	}
}

// UIcontroller is a task for user interface handling
func UIcontroller() {

	ui := UIctrl{}
	ui.Init()

	ui.Scr.SendToDisplay()
	ui.Scr.UpdateSize()

	inp := bufio.NewReader(os.Stdin)
	for {

		switch r, _, _ := inp.ReadRune(); r {
		case '\x1b': // Esc
			ui.DeInit()
			os.Exit(0)
		case '\t': // Tab

			if ui.CommonBox.IsActive {
				ui.CommonBox.SetActiveState(false)
				ui.TableBox.SetActiveState(true)
			} else {
				ui.CommonBox.SetActiveState(true)
				ui.TableBox.SetActiveState(false)
			}

			ui.debugPrintf("cmnbox state: %v", ui.CommonBox.IsActive)

			ui.Draw(&ui.Scr)
			ui.Scr.SendToDisplay()
		default:
			ui.Scr.SendToDisplay()
		}
	}

	for {

	}
}
