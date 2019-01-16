package ui

import "github.com/phpinfo/gosnake/renderer"

type LineHorizontal struct {
	x1, x2, y int
	ch rune
}

func NewLineHorizontal(x1, x2, y int, ch rune) *LineHorizontal {
	return &LineHorizontal{x1, x2, y, ch}
}

func (line *LineHorizontal) Render(renderer renderer.Renderer) {
	for x := line.x1; x < line.x2; x++ {
		renderer.Cell(x, line.y, line.ch)
	}
}

func (line *LineHorizontal) GenerateCells() []Cell {
	result := make([]Cell, line.x2 - line.x1, line.x2 - line.x1)

	for x := line.x1; x < line.x2; x++ {
		result = append(result, *NewCell(x, line.y, line.ch))
	}

	return result
}