package geometry

type Point struct {
	X, Y int
}

func NewPoint(x, y int) *Point {
	return &Point{X: x, Y: y}
}

func (point *Point) Add(x, y int) *Point {
	return NewPoint(point.X + x, point.Y + y)
}

func (point *Point) Equals(point2 *Point) bool {
	return point.X == point2.X && point.Y == point2.Y
}
