package main

import (
	"bufio"
	"flydb/cui"
	"fmt"
	"os"
)

// Used symbols ┌ ─ ┐ └ │ ┘
func main() {
	cyclesCnt := 0
	var scr cui.Screen
	scr.Init()

	reader := bufio.NewReader(os.Stdin)

	for {
		scr.UpdateSize()
		width := scr.GetWidth()
		height := scr.GetHeight()

		scr.SetRune('┌', 0, 0)
		scr.SetRune('┐', width-1, 0)
		scr.SetRune('└', 0, height-1)
		scr.SetRune('┘', width-1, height-1)

		scr.SendToDisplay()

		_, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
			return
		} // if result == '1'

		cyclesCnt++
		//scr.Clear()
	}

	reader.ReadRune()
}
