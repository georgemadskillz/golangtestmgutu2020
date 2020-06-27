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

	keyInfo := InfoBox{1, 1, 25, 7, false, nil}
	keyInfo.Init()
	keyInfo.SetLineText(0, "Shortcuts")
	keyInfo.SetLineText(1, "Esc: quit")
	keyInfo.SetLineText(2, "Tab: next window")
	keyInfo.Draw(&scr)

	testInfo := InfoBox{30, 15, 25, 15, false, nil}
	testInfo.Init()
	testInfo.SetLineText(0, "Hello, world!")
	testInfo.SetLineText(1, "blah blah ")
	testInfo.SetLineText(8, "_____ ** ___")
	testInfo.Draw(&scr)

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
