package snake

import (
	"github.com/nsf/termbox-go"
	"github.com/phpinfo/gosnake/geometry"
)

func cell(point *geometry.Point, ch rune) {
	termbox.SetCell(point.X, point.Y, ch, termbox.ColorDefault, termbox.ColorDefault)
}

func hline(p1, p2 *geometry.Point, ch rune) {
	for x := p1.X; x < p2.X; x++ {
		cell(geometry.NewPoint(x, p1.Y), ch)
	}
}

func vline(p1, p2 *geometry.Point, ch rune) {
	for y := p1.Y; y < p2.Y; y++ {
		cell(geometry.NewPoint(p1.X, y), ch)
	}
}
