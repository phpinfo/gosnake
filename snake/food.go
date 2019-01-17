package snake

import (
	"github.com/phpinfo/gosnake/geometry"
)

type Food struct {
	Point *geometry.Point
}

func NewFood(point *geometry.Point) *Food {
	return &Food{point}
}
