package ui

import "github.com/phpinfo/gosnake/renderer"

type LineVertical struct {
	x, y1, y2 int
	ch rune
}

func NewLineVertical(x, y1, y2 int, ch rune) *LineVertical {
	return &LineVertical{x, y1, y2, ch}
}

func (line *LineVertical) Render(renderer renderer.Renderer) {
	for y := line.y1; y < line.y2; y++ {
		renderer.Cell(line.x, y, line.ch)
	}
}

func (line *LineVertical) GenerateCells() []Cell {
	result := make([]Cell, line.y2 - line.y1, line.y2 - line.y1)

	for y := line.y1; y < line.y2; y++ {
		result = append(result, *NewCell(line.x, y, line.ch))
	}

	return result
}
