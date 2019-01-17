package ui

import "github.com/phpinfo/gosnake/geometry"

const (
	CharFood = '@'
)

type Food struct {
	point *geometry.Point
}

func NewFood(point *geometry.Point) *Food {
	return &Food{point}
}

func (food *Food) GenerateCells() []Cell {
	cell := NewCell(food.point.X, food.point.Y, CharFood)
	return []Cell{*cell}
}
