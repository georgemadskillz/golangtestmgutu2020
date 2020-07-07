package cui

// UIctrl is a common structure handling UI
type UIctrl struct {
	Scr       Screen
	Extbox    Box
	CommonBox InfoBox
	StatusBox InfoBox
	TblBox    TableBox
}

// Init initialize UI controller
func (ui *UIctrl) Init() {

	ui.Scr.Init()

	// Ext border
	ui.Extbox = Box{0, 0, ui.Scr.Width, ui.Scr.Height, false}

	// Common window
	ui.CommonBox = InfoBox{1, 1, 40, 10, false, nil}
	ui.CommonBox.Init()
	ui.CommonBox.SetLineText(0, "БАЗА ДАННЫХ")
	ui.CommonBox.SetLineText(2, "Выбор таблицы:")
	ui.CommonBox.SetLineText(4, "# Рейсы")
	ui.CommonBox.SetLineText(5, "# Аэропорты")
	ui.CommonBox.SetLineText(6, "# Цены")
	ui.CommonBox.SetActiveState(true)

	// Table window
	ui.TblBox = TableBox{41, 1, ui.Scr.Width - 2 - 40, ui.Scr.Height - 2, false, 4, nil, nil}
	ui.TblBox.Init()
	ui.TblBox.SetCell(0, 0, "Date from")
	ui.TblBox.SetCell(0, 1, "From")
	ui.TblBox.SetCell(0, 2, "To")
	ui.TblBox.SetCell(0, 3, "Date to")

	ui.TblBox.FillCell(1, 0, '═')
	ui.TblBox.FillCell(1, 1, '═')
	ui.TblBox.FillCell(1, 2, '═')
	ui.TblBox.FillCell(1, 3, '═')

	// Status window
	ui.StatusBox = InfoBox{1, 11, 40, ui.Scr.Height - 2 - 10, false, nil}
	ui.StatusBox.Init()
	ui.StatusBox.SetLineText(0, "Статус программы:")

	ui.StatusBox.SetLineText(2, "Текущая таблица: <Рейсы>")

	ui.StatusBox.SetLineText(14, "Управление программой:")
	ui.StatusBox.SetLineText(15, "Esc: выход из программы")
	ui.StatusBox.SetLineText(16, "Tab: переход между окнами")

	ui.StatusBox.SetLineText(18, "─────────────────────────────────────")
	ui.StatusBox.SetLineText(19, "Отладочная информация:")

	ui.Scr.UpdateSize()
	ui.Draw(&ui.Scr)
	ui.Scr.SendToDisplay()
}

// DeInit is
func (ui *UIctrl) DeInit() {
	ui.Scr.Clear()
}

// Draw draws all it's elements
func (ui *UIctrl) Draw(scr *Screen) {
	ui.Extbox.Draw(scr)
	ui.CommonBox.Draw(scr)
	ui.StatusBox.Draw(scr)
	ui.TblBox.Draw(scr)
}
