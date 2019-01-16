package ui

import (
	"fmt"
)

type Counter struct {
	value    *int
	x, y     int
	template string
}

func NewCounter(value *int, x, y int) *Counter {
	return &Counter{
		value: value,
		x: x,
		y: y,
		template: "%03d",
	}
}

func (counter *Counter) GenerateCells() []Cell {
	text := fmt.Sprintf(counter.template, *counter.value)
	result := make([]Cell, len(text), len(text))

	for dx, ch := range text {
		result = append(result, *NewCell(counter.x + dx, counter.y, ch))
	}

	return result
}
