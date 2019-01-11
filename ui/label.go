package ui

type Label struct {
	text string
	x, y int
}

func NewLabel (text string, x, y int) *Label {
	return &Label{text, x, y}
}

func (label *Label) GenerateCells() []Cell {
	cells := make([]Cell, len(label.text), len(label.text))
	for dx, ch := range label.text {
		cells[dx] = *NewCell(label.x + dx, label.y, ch)
	}

	return cells
}
