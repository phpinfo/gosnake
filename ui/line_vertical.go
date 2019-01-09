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
