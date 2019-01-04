package snake

const (
	DirectionUp   Direction = 1 + iota
	DirectionDown
	DirectionLeft
	DirectionRight
)

type Direction uint8

func (direction Direction) GetMovement() (int, int) {
	switch direction {
	case DirectionUp:
		return 0, -1
	case DirectionDown:
		return 0, 1
	case DirectionRight:
		return 1, 0
	case DirectionLeft:
		return -1, 0
	}
	return 0, 0
}

func (direction Direction) IsOpposite(direction2 Direction) bool {
	oppositeDirections := map[Direction]Direction{
		DirectionLeft:  DirectionRight,
		DirectionRight: DirectionLeft,
		DirectionUp:    DirectionDown,
		DirectionDown:  DirectionUp,
	}

	return oppositeDirections[direction] == direction2
}
