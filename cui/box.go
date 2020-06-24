package cui

// Box is a rectangle with width*height at (x,y)
type Box struct {
	X, Y          int
	Width, Height int
	IsActive      bool
}

// // New initializes new box with given coords and size
// func (b *Box) New(x, y, w, h int) *Box {
// 	return &Box{x, y, w, h, false}
// }

// SetSize sets new size of the box
func (b *Box) SetSize(w, h int) {
	b.Width = w
	b.Height = h
}

// SetPosition sets new position of the box
func (b *Box) SetPosition(x, y int) {
	b.X = x
	b.Y = y
}

// SetActiveState set active/not active state of the box
func (b *Box) SetActiveState(state bool) {
	b.IsActive = state
}

// Draw draws empty box at screen buffer
// and fill its internal space with space symbols
func (b *Box) Draw(scr *Screen) {
	for x := b.X + 1; x < b.X+b.Width-1; x++ {
		for y := b.Y + 1; y < b.Y+b.Height-1; y++ {
			scr.SetRune(x, y, ' ')
		}
	}

	for x := b.X + 1; x < b.X+b.Width-1; x++ {
		scr.SetRune(x, b.Y, '─')
		scr.SetRune(x, b.Y+b.Height-1, '─')
	}

	for y := b.Y + 1; y < b.Y+b.Height-1; y++ {
		scr.SetRune(b.X, y, '│')
		scr.SetRune(b.X+b.Width-1, y, '│')
	}

	scr.SetRune(b.X, b.Y, '┌')
	scr.SetRune(b.X+b.Width-1, b.Y, '┐')
	scr.SetRune(b.X, b.Y+b.Height-1, '└')
	scr.SetRune(b.X+b.Width-1, b.Y+b.Height-1, '┘')
}
