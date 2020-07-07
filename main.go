package main

import (
	"bufio"
	"os"

	"flydb/appctrl"
)

func main() {
	app := appctrl.AppController{}
	app.Init()

	inp := bufio.NewReader(os.Stdin)
	for {

		kbdKey, _, _ := inp.ReadRune()

		app.CheckKbdInput(kbdKey)
	}
}
