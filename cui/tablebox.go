package cui

// TableBox is
type TableBox struct {
	X, Y          int
	Width, Height int
	IsActive      bool
	ColsAmount    int
	Cells         [][]string
	CellsWidths   []int
}

// Init creates
func (t *TableBox) Init() {

	rows := t.Height - 2
	for r := 0; r < rows; r++ {

		row := make([]string, t.ColsAmount, t.ColsAmount)
		t.Cells = append(t.Cells, row)
	}

	t.CellsWidths = make([]int, t.ColsAmount, t.ColsAmount)

	w := t.Width / t.ColsAmount
	for i := range t.CellsWidths {
		if i == len(t.CellsWidths)-1 {
			t.CellsWidths[i] = w + t.Width%t.ColsAmount
		} else {
			t.CellsWidths[i] = w
		}
	}

	for r := 0; r < t.Height-2; r++ {
		for c := 0; c < t.ColsAmount; c++ {
			t.SetCell(r, c, " ")
		}
	}
}

// SetColsWidths is
func (t *TableBox) SetColsWidths(widths ...int) {
	if len(widths) != t.ColsAmount {
		return
	}

}

// SetCell is
func (t *TableBox) SetCell(row, col int, text string) {
	if row < 0 || row >= t.Height-2 {
		return
	}

	if col < 0 || col >= t.ColsAmount {
		return
	}

	cellStr := ""

	if col > 0 {
		cellStr += "║"
	}

	cellStr += text

	runes := []rune(cellStr)

	if len(runes) >= t.CellsWidths[col] {
		t.Cells[row][col] = string(runes[0:t.CellsWidths[col]])
	} else {
		for len(runes) < t.CellsWidths[col] {
			runes = append(runes, ' ')
		}
		t.Cells[row][col] = string(runes)
	}
}

// FillCell is
func (t *TableBox) FillCell(row, col int, r rune) {
	str := ""
	for i := 0; i < t.CellsWidths[col]; i++ {
		str += string(r)
	}

	t.SetCell(row, col, str)
}

// SetSize sets new size of the infobox
func (t *TableBox) SetSize(w, h int) {
	t.Width = w
	t.Height = h
}

// SetPosition sets new position of the infobox
func (t *TableBox) SetPosition(x, y int) {
	t.X = x
	t.Y = y
}

// SetActiveState set active/not active state of the infobox
func (t *TableBox) SetActiveState(state bool) {
	t.IsActive = state
}

// GetActiveState returns IsActive of UI element
func (t *TableBox) GetActiveState() bool {
	return t.IsActive
}

// Clean is
func (t *TableBox) Clean(scr *Screen) {
	for x := t.X + 1; x < t.X+t.Width-1; x++ {
		for y := t.Y + 1; y < t.Y+t.Height-1; y++ {
			scr.SetRune(x, y, ' ')
		}
	}
}

// Draw is
func (t *TableBox) Draw(scr *Screen) {
	t.Clean(scr)

	for r, row := range t.Cells {
		rowStr := ""
		for _, cellVal := range row {
			rowStr += cellVal
		}

		strRunes := []rune(rowStr)
		y := t.Y + 1 + r
		j := 0
		for x := t.X + 1; x < t.X+t.Width-1; x++ {
			scr.SetRune(x, y, strRunes[j])
			r++
			j++
			if j >= len(strRunes) {
				break
			}
		}
	}

	for x := t.X + 1; x < t.X+t.Width-1; x++ {
		if t.IsActive {
			scr.SetRune(x, t.Y, '━')
			scr.SetRune(x, t.Y+t.Height-1, '━')
			// scr.SetRune(x, b.Y, '═')
			// scr.SetRune(x, b.Y+b.Height-1, '═')
		} else {
			scr.SetRune(x, t.Y, '─')
			scr.SetRune(x, t.Y+t.Height-1, '─')
		}

	}

	for y := t.Y + 1; y < t.Y+t.Height-1; y++ {
		if t.IsActive {
			scr.SetRune(t.X, y, '┃')
			scr.SetRune(t.X+t.Width-1, y, '┃')
			// scr.SetRune(b.X, y, '║')
			// scr.SetRune(b.X+b.Width-1, y, '║')
		} else {
			scr.SetRune(t.X, y, '│')
			scr.SetRune(t.X+t.Width-1, y, '│')
		}

	}

	if t.IsActive {
		scr.SetRune(t.X, t.Y, '┏')
		scr.SetRune(t.X+t.Width-1, t.Y, '┓')
		scr.SetRune(t.X, t.Y+t.Height-1, '┗')
		scr.SetRune(t.X+t.Width-1, t.Y+t.Height-1, '┛')
		// scr.SetRune(b.X, b.Y, '╔')
		// scr.SetRune(b.X+b.Width-1, b.Y, '╗')
		// scr.SetRune(b.X, b.Y+b.Height-1, '╚')
		// scr.SetRune(b.X+b.Width-1, b.Y+b.Height-1, '╝')
	} else {
		scr.SetRune(t.X, t.Y, '┌')
		scr.SetRune(t.X+t.Width-1, t.Y, '┐')
		scr.SetRune(t.X, t.Y+t.Height-1, '└')
		scr.SetRune(t.X+t.Width-1, t.Y+t.Height-1, '┘')
	}

}
