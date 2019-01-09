package snake

import (
	"github.com/phpinfo/gosnake/geometry"
	"github.com/phpinfo/gosnake/renderer"
)

type Snake struct {
	Body      []*geometry.Point
	Direction Direction
}

func NewSnake(body []*geometry.Point, direction Direction) *Snake {
	return &Snake{body, direction}
}

func (snake *Snake) Head() *geometry.Point {
	return snake.Body[len(snake.Body) - 1]
}

func (snake *Snake) SetDirection(direction Direction) {
	if snake.Direction.IsOpposite(direction) {
		return
	}
	snake.Direction = direction
}

func (snake *Snake) Move() {
	dx, dy := snake.Direction.GetMovement()
	head := snake.Head().Add(dx, dy)
	snake.Body = append(snake.Body[1:], head)
}

func (snake *Snake) Eat() {
	dx, dy := snake.Direction.GetMovement()
	head := snake.Head().Add(dx, dy)
	snake.Body = append(snake.Body, head)
}

func (snake *Snake) Contains(point *geometry.Point) bool {
	return snake.contains(point, snake.Body)
}

func (snake *Snake) SelfCollides() bool {
	return snake.contains(snake.Head(), snake.Body[:len(snake.Body)-1])
}

func (snake *Snake) contains(point *geometry.Point, body []*geometry.Point) bool {
	for _, p := range body {
		if p.Equals(point) {
			return true
		}
	}
	return false
}

func (snake *Snake) Render(renderer renderer.Renderer) {
	for _, point := range snake.Body {
		renderer.Cell(point.X, point.Y, '*')
	}
}
