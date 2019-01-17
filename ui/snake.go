package ui

import "github.com/phpinfo/gosnake/geometry"

const (
	CharBody = '#'
)

type Snake struct {
	points []*geometry.Point
}

func NewPointCollection(points []*geometry.Point) *Snake {
	return &Snake{points: points}
}

func (pointCollection *Snake) SetPoints(points []*geometry.Point) {
	pointCollection.points = points
}

func (pointCollection *Snake) GenerateCells() []Cell {
	result := make([]Cell, len(pointCollection.points), len(pointCollection.points))

	for _, point := range pointCollection.points {
		result = append(result, *NewCell(point.X, point.Y, CharBody))
	}

	return result
}
