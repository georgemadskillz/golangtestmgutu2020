package cui

import "fmt"

// Painter type implements painting to Screen at coordinates
type Painter struct {
	Scr    *Screen
	canvas [][]rune
}

// Init painter
func (p *Painter) Init(scr *Screen) {
	fmt.Printf("Painter init..\n")

	height := int(scr.GetHeight())
	fmt.Printf("Painter init: window height is %v..\n", height)
	width := int(scr.GetWidth())
	fmt.Printf("Painter init: window width is %v..\n", width)

	p.canvas = make([][]rune, width)

	for row := 0; row < height; row++ {
		//fmt.Printf("Painter init: row=%v..\n", row)
		p.canvas[row] = make([]rune, width)
		//fmt.Printf("Painter init: row[%v]: %v..\n", row, p.canvas[row])
	}
}

// DrawRune draws one symbol at given coords
func (p *Painter) DrawRune(w int, h int, r rune) {
	//fmt.Printf("Painter: try draw rune %v at x=%v, y=%v..\n", r, w, h)
	p.canvas[h][w] = r
	//fmt.Printf("Painter: ok..\n")
}

// SendToScreen sends painter buffer into screen buffer
func (p *Painter) SendToScreen(scr *Screen) {
	fmt.Printf("Sned to screen..\n")

	buff := scr.GetBuff()

	buff.Reset()
	for i := range p.canvas {
		str := string(p.canvas[i]) + "\n"
		fmt.Printf("Apend: %v\n", str)
		buff.WriteString(str)
	}
}
