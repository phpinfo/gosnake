package snake

import (
	"fmt"
	"github.com/phpinfo/gosnake/geometry"
	"github.com/phpinfo/gosnake/renderer"
)

type Counter struct {
	value    *int
	point    *geometry.Point
	template string
}

func NewCounter(value *int, point *geometry.Point) *Counter {
	return &Counter{value, point, "%03d"}
}

func (counter *Counter) Render (renderer renderer.Renderer) {
	text := fmt.Sprintf(counter.template, *counter.value)

	for dx, ch := range text {
		point := counter.point.Add(dx, 0)
		renderer.Cell(point.X, point.Y, ch)
	}
}
