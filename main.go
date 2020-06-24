package main

import (
	"bufio"
	"flydb/cui"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

/*
Roadmap
1) интерфейс консоли
2) модуль хранения данных в оперативке
3) файловый вод-вывод

Fixes:
! брать размеры терминала с помощью пакета terminal, пока что делаем вручную системными вызовами

*/

func main() {
	var scr cui.Screen
	scr.Init()

	go userInpController(&scr)

	scrW := scr.GetWidth()
	scrH := scr.GetHeight()

	screenExtbox := cui.Box{0, 0, scrW, scrH, false}
	screenExtbox.Draw(&scr)

	screenBox1 := cui.Box{1, 1, 5, 10, false}
	screenBox1.Draw(&scr)

	screenBox2 := cui.Box{15, 1, 5, 10, false}
	screenBox2.Draw(&scr)

	screenBox3 := cui.Box{30, 1, 5, 10, false}
	screenBox3.Draw(&scr)

	screenBox4 := cui.Box{1, 30, 10, 5, false}
	screenBox4.Draw(&scr)

	screenBox5 := cui.Box{50, 20, 15, 15, false}
	screenBox5.Draw(&scr)

	for {

	}
}

func userInpController(scr *cui.Screen) {

	state, err := terminal.MakeRaw(0)
	if err != nil {
		//log.Fatalln("setting stdin to raw:", err)
	}

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
}
