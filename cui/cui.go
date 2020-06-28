package cui

import (
	"bufio"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// UIctrl is a common structure handling UI
type UIctrl struct {
	Scr       Screen
	TermState *terminal.State
	Elements  []UIelement
}

// UIelement is
type UIelement struct {
	ID  int
	Obj interface{}
}

// Init initialize UI controller
func (ui *UIctrl) Init() {
	ui.Scr.Init()

	state, err := terminal.MakeRaw(0)
	if err != nil {
		//log.Fatalln("setting stdin to raw:", err)
	}
	ui.TermState = state

	// Init user inteface
	ui.Elements = make([]UIelement, 3)

	// Ext border
	screenExtbox := Box{0, 0, ui.Scr.Width, ui.Scr.Height, false}
	screenExtbox.Draw(&ui.Scr)
	ui.Elements = append(ui.Elements, UIelement{0, &screenExtbox})

	// Common window
	commonBox := InfoBox{1, 1, 40, 10, false, nil}
	commonBox.Init()
	commonBox.SetLineText(0, "БАЗА ДАННЫХ")
	commonBox.SetLineText(2, "Выбор таблицы:")
	commonBox.SetLineText(4, "# Рейсы")
	commonBox.SetLineText(5, "# Аэропорты")
	commonBox.SetLineText(6, "# Цены")
	commonBox.Draw(&ui.Scr)
	ui.Elements = append(ui.Elements, UIelement{1, &commonBox})

	// Table window
	tableBox := InfoBox{41, 1, ui.Scr.Width - 2 - 40, ui.Scr.Height - 2, false, nil}
	tableBox.Init()
	tableBox.SetLineText(0, "Здесь будет таблица базы данных")
	tableBox.Draw(&ui.Scr)
	ui.Elements = append(ui.Elements, UIelement{1, &tableBox})

	// Status window
	statusBox := InfoBox{1, 11, 40, ui.Scr.Height - 2 - 10, false, nil}
	statusBox.Init()
	statusBox.SetLineText(0, "Статус программы:")

	statusBox.SetLineText(2, "Текущая таблица: <Рейсы>")

	statusBox.SetLineText(20, "Управление программой:")
	statusBox.SetLineText(21, "Esc: выход из программы")
	statusBox.SetLineText(22, "Tab: переход между окнами")
	statusBox.Draw(&ui.Scr)
	ui.Elements = append(ui.Elements, UIelement{1, &statusBox})
}

// DeInit restores terminal state from raw and clears screen
func (ui *UIctrl) DeInit() {
	if err := terminal.Restore(0, ui.TermState); err != nil {
		//log.Println("warning, failed to restore terminal:", err)
	}
	ui.Scr.Clear()
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
		case '\x1b':
			ui.DeInit()
			os.Exit(0)
		case '\t': // Tab

		default:
			ui.Scr.SendToDisplay()
		}
	}

	for {

	}
}
