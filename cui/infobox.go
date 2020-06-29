package cui

// InfoBox is a rectangle with lines of text
type InfoBox struct {
	X, Y          int
	Width, Height int
	IsActive      bool
	Text          []string
}

// Init creates new box with given coords and size,
// and make text slice
func (b *InfoBox) Init() {

	N := b.Height - 2
	text := make([]string, N)

	for i := 0; i < N; i++ {
		text = append(text, "")
	}

	b.Text = text
}

// SetLineText set up given text into given line of infobox
func (b *InfoBox) SetLineText(line int, text string) {
	if line < 0 || line >= b.Height-2 {
		return
	}

	b.Text[line] = text
}

// SetSize sets new size of the infobox
func (b *InfoBox) SetSize(w, h int) {
	b.Width = w
	b.Height = h
}

// SetPosition sets new position of the infobox
func (b *InfoBox) SetPosition(x, y int) {
	b.X = x
	b.Y = y
}

// SetActiveState set active/not active state of the infobox
func (b *InfoBox) SetActiveState(state bool) {
	b.IsActive = state
}

// GetActiveState returns IsActive of UI element
func (b *InfoBox) GetActiveState() bool {
	return b.IsActive
}

// Clean clear infobox with spaces within borders
func (b *InfoBox) Clean(scr *Screen) {
	for x := b.X + 1; x < b.X+b.Width-1; x++ {
		for y := b.Y + 1; y < b.Y+b.Height-1; y++ {
			scr.SetRune(x, y, ' ')
		}
	}
}

// Draw draws empty infobox at screen buffer
// and fill its internal space with space symbols
func (b *InfoBox) Draw(scr *Screen) {
	b.Clean(scr)

	//fmt.Printf("Infobox: %+v\r\n", b)
	//fmt.Printf("Draw infobox.. \r\n")
	for i := range b.Text {
		//fmt.Printf("i=%v: text[i]=|%v|\r\n", i, b.Text[i])
		lineRunes := []rune(b.Text[i])

		if len(lineRunes) == 0 {
			continue
		}

		y := b.Y + 1 + i
		j := 0
		for x := b.X + 1; x < b.X+b.Width-1; x++ {
			//fmt.Printf("j=%v: Set rune %q at x=%v,y=%v\r\n", j, lineRunes[j], x, y)
			scr.SetRune(x, y, lineRunes[j])
			i++
			j++
			if j >= len(lineRunes) {
				break
			}
		}
	}

	for x := b.X + 1; x < b.X+b.Width-1; x++ {
		if b.IsActive {
			scr.SetRune(x, b.Y, '━')
			scr.SetRune(x, b.Y+b.Height-1, '━')
			// scr.SetRune(x, b.Y, '═')
			// scr.SetRune(x, b.Y+b.Height-1, '═')
		} else {
			scr.SetRune(x, b.Y, '─')
			scr.SetRune(x, b.Y+b.Height-1, '─')
		}

	}

	for y := b.Y + 1; y < b.Y+b.Height-1; y++ {
		if b.IsActive {
			scr.SetRune(b.X, y, '┃')
			scr.SetRune(b.X+b.Width-1, y, '┃')
			// scr.SetRune(b.X, y, '║')
			// scr.SetRune(b.X+b.Width-1, y, '║')
		} else {
			scr.SetRune(b.X, y, '│')
			scr.SetRune(b.X+b.Width-1, y, '│')
		}

	}

	if b.IsActive {
		scr.SetRune(b.X, b.Y, '┏')
		scr.SetRune(b.X+b.Width-1, b.Y, '┓')
		scr.SetRune(b.X, b.Y+b.Height-1, '┗')
		scr.SetRune(b.X+b.Width-1, b.Y+b.Height-1, '┛')
		// scr.SetRune(b.X, b.Y, '╔')
		// scr.SetRune(b.X+b.Width-1, b.Y, '╗')
		// scr.SetRune(b.X, b.Y+b.Height-1, '╚')
		// scr.SetRune(b.X+b.Width-1, b.Y+b.Height-1, '╝')
	} else {
		scr.SetRune(b.X, b.Y, '┌')
		scr.SetRune(b.X+b.Width-1, b.Y, '┐')
		scr.SetRune(b.X, b.Y+b.Height-1, '└')
		scr.SetRune(b.X+b.Width-1, b.Y+b.Height-1, '┘')
	}

}
