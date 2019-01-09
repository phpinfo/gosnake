package snake

import (
	"github.com/phpinfo/gosnake/geometry"
	"github.com/phpinfo/gosnake/renderer"
)

type Food struct {
	Point *geometry.Point
}

func NewFood(point *geometry.Point) *Food {
	return &Food{point}
}

func (food *Food) Render(renderer renderer.Renderer) {
	renderer.Cell(food.Point.X, food.Point.Y, '@')
}
