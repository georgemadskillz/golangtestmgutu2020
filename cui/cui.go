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

	screenBox1 := Box{1, 1, 5, 10, false}
	screenBox1.Draw(&scr)

	screenBox2 := Box{15, 1, 5, 10, false}
	screenBox2.Draw(&scr)

	screenBox3 := Box{30, 1, 5, 10, false}
	screenBox3.Draw(&scr)

	screenBox4 := Box{1, 30, 10, 5, false}
	screenBox4.Draw(&scr)

	screenBox5 := Box{50, 20, 15, 15, false}
	screenBox5.IsActive = true
	screenBox5.Draw(&scr)

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
