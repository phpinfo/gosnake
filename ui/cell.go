package ui

type Cell struct {
	X, Y int
	Ch rune
}

func NewCell(x, y int, ch rune) *Cell {
	return &Cell{X: x, Y: y, Ch: ch}
}
