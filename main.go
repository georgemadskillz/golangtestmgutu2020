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

	go func(scr *cui.Screen) {

		state, err := terminal.MakeRaw(0)
		if err != nil {
			//log.Fatalln("setting stdin to raw:", err)
		}

		scr.SendToDisplay()

		inp := bufio.NewReader(os.Stdin)
		for {

			switch r, _, _ := inp.ReadRune(); r {
			case '\x1b':
				if err := terminal.Restore(0, state); err != nil {
					//log.Println("warning, failed to restore terminal:", err)
				}
				os.Exit(0)
			default:
				scr.SendToDisplay()
			}
		}
	}(&scr)

	for {
		scr.UpdateSize()
		width := scr.GetWidth()
		height := scr.GetHeight()

		scr.SetRune('┌', 0, 0)
		scr.SetRune('┐', width-1, 0)
		scr.SetRune('└', 0, height-1)
		scr.SetRune('┘', width-1, height-1)

	}
}
