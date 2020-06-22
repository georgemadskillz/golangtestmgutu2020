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
	var scr cui.creen
	var paint cui.Painter

	scr.Init()
	paint.Init(&scr)

	reader := bufio.NewReader(os.Stdin)

	for {
		cols := int(scr.GetWidth())
		rows := int(scr.GetHeight())

		var c int
		var r int
		for r = 0; c < rows; c++ {
			for c = 0; r < cols; r++ {
				paint.DrawRune(c, r, rune(c))
			}

			paint.DrawRune(c, r, '\n')
		}

		paint.SendToScreen(&scr)
		scr.Draw()

		_, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
			return
		} // if result == '1'

		cyclesCnt++
		//scr.Clear()
	}
}
