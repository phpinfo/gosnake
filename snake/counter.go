package snake

import (
	"fmt"
	"github.com/phpinfo/gosnake/geometry"
)

type Counter struct {
	value    *int
	point    *geometry.Point
	template string
}

func NewCounter(value *int, point *geometry.Point) *Counter {
	return &Counter{value, point, "%03d"}
}

func (counter *Counter) Render () {
	text := fmt.Sprintf(counter.template, *counter.value)

	for dx, c := range text {
		cell(counter.point.Add(dx, 0), c)
	}
}
