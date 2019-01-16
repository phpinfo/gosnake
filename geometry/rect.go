package geometry

type Rect struct {
	LeftTopPoint, RightTopPoint, LeftBottomPoint, RightBottomPoint *Point
	Width, Height, Left, Top, Right, Bottom int
}

func NewRect(x, y, w, h int) *Rect {
	var (
		left   = x
		top    = y
		right  = x + w
		bottom = y + h
	)

	return &Rect{
		LeftTopPoint:     NewPoint(left, top),
		RightTopPoint:    NewPoint(right, top),
		LeftBottomPoint:  NewPoint(left, bottom),
		RightBottomPoint: NewPoint(right, bottom),
		Width:            w,
		Height:           h,
		Left:             left,
		Top:              top,
		Right:            right,
		Bottom:           bottom,
	}
}

func (rect *Rect) Contains(point Point) bool {
	return point.X >= rect.Left &&
		point.X < rect.Right &&
		point.Y >= rect.Top &&
		point.Y < rect.Bottom
}
