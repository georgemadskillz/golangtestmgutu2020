package cui

import (
	"bufio"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// UIcontroller is a task for user interface handling
func UIcontroller() {

	var scr Screen
	scr.Init()

	state, err := terminal.MakeRaw(0)
	if err != nil {
		//log.Fatalln("setting stdin to raw:", err)
	}

	scrW := scr.GetWidth()
	scrH := scr.GetHeight()

	screenExtbox := Box{0, 0, scrW, scrH, false}
	screenExtbox.Draw(&scr)

	commonBox := InfoBox{1, 1, 40, 10, false, nil}
	commonBox.Init()
	commonBox.SetLineText(0, "БАЗА ДАННЫХ")
	commonBox.SetLineText(2, "Выбор таблицы:")
	commonBox.SetLineText(4, "# Рейсы")
	commonBox.SetLineText(5, "# Аэропорты")
	commonBox.SetLineText(6, "# Цены")
	commonBox.Draw(&scr)

	tableBox := InfoBox{41, 1, scrW - 2 - 40, scrH - 2, false, nil}
	tableBox.Init()
	tableBox.SetLineText(0, "Здесь будет таблица базы данных")
	tableBox.Draw(&scr)

	statusBox := InfoBox{1, 11, 40, scrH - 2 - 10, false, nil}
	statusBox.Init()
	statusBox.SetLineText(0, "Статус программы:")

	statusBox.SetLineText(2, "Текущая таблица: <Рейсы>")

	statusBox.SetLineText(20, "Управление программой:")
	statusBox.SetLineText(21, "Esc: выход из программы")
	statusBox.SetLineText(22, "Tab: переход между окнами")
	statusBox.Draw(&scr)

	scr.SendToDisplay()
	scr.UpdateSize()

	inp := bufio.NewReader(os.Stdin)
	for {

		switch r, _, _ := inp.ReadRune(); r {
		case '\x1b':
			if err := terminal.Restore(0, state); err != nil {
				//log.Println("warning, failed to restore terminal:", err)
			}
			scr.Clear()
			os.Exit(0)
		case '\t': // Tab

		default:
			scr.SendToDisplay()
		}
	}

	for {

	}
}
