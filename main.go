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
	var paint cui.Painter

	scr.Init()
	paint.Init(&scr)

	reader := bufio.NewReader(os.Stdin)

	for {
		width := int(scr.GetWidth())
		height := int(scr.GetHeight())

		paint.DrawRune(0, 0, '┌')
		paint.DrawRune(width-1, 0, '┐')
		paint.DrawRune(0, height-1, '└')
		paint.DrawRune(width-1, height-1, '┘')

		cnt := 0
		for w := 0; w < width-1; w++ {
			for h := 0; h < height-1; h++ {
				paint.DrawRune(w, 0, rune(cnt))
				cnt++
			}
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
