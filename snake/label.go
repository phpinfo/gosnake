package snake

import "github.com/phpinfo/gosnake/geometry"

type Label struct {
	text string
	point *geometry.Point
}

func NewLabel (text string, point *geometry.Point) *Label {
	return &Label{text, point}
}

func (label *Label) Render () {
	for dx, c := range label.text {
		cell(label.point.Add(dx, 0), c)
	}
}
