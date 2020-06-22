package cui

// Painter type implements painting to Screen at coordinates
type Painter struct {
	Scr    *Screen
	canvas [][]rune
}

// Init painter
func (p *Painter) Init(scr *Screen) {
	w := scr.GetWidth()
	h := scr.GetHeight()

	if w > h {
		p.canvas = make([][]rune, w)
	} else {
		p.canvas = make([][]rune, h)
	}
}

// DrawRune draws one symbol at given coords
func (p *Painter) DrawRune(x int, y int, r rune) {
	p.canvas[x][y] = r
}

// SendToScreen sends painter buffer into screen buffer
func (p *Painter) SendToScreen(scr *Screen) {
	buff := scr.GetBuff()

	buff.Reset()
	for i := range p.canvas {
		buff.WriteString(string(p.canvas[i]) + "\n")
	}
}
