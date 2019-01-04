package snake

type Box struct {
	Rect *Rect
	Snake *Snake
}

func NewBox(rect *Rect, snake *Snake) *Box {
	return &Box{rect, snake}
}
