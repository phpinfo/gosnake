package snake

import "fmt"

type Counter struct {
	value    *int
	point    *Point
	template string
}

func NewCounter(value *int, point *Point) *Counter {
	return &Counter{value, point, "%03d"}
}

func (counter *Counter) Render () {
	text := fmt.Sprintf(counter.template, *counter.value)

	for dx, c := range text {
		cell(counter.point.Add(dx, 0), c)
	}
}
