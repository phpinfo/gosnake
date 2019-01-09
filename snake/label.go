package snake

import (
	"github.com/phpinfo/gosnake/geometry"
	"github.com/phpinfo/gosnake/renderer"
)

type Label struct {
	text string
	point *geometry.Point
}

func NewLabel (text string, point *geometry.Point) *Label {
	return &Label{text, point}
}

func (label *Label) Render (renderer renderer.Renderer) {
	for dx, ch := range label.text {
		point := label.point.Add(dx, 0)
		renderer.Cell(point.X, point.Y, ch)
	}
}
