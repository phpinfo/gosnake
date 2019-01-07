package snake

import "github.com/phpinfo/gosnake/geometry"

const (
	CharCornerLeftTop     = '╔'
	CharCornerLeftBottom  = '╚'
	CharCornerRightTop    = '╗'
	CharCornerRightBottom = '╝'
	CharHorizontal        = '═'
	CharVertical          = '║'
)

type Rect struct {
	LeftTopPoint, RightTopPoint, LeftBottomPoint, RightBottomPoint *geometry.Point
	Dimensions *geometry.Dimensions
	Left, Top, Right, Bottom int
}

func NewRect(x, y, w, h int) *Rect {
	var (
		left   = x
		top    = y
		right  = x + w
		bottom = y + h
	)

	return &Rect{
		LeftTopPoint:     geometry.NewPoint(left, top),
		RightTopPoint:    geometry.NewPoint(right, top),
		LeftBottomPoint:  geometry.NewPoint(left, bottom),
		RightBottomPoint: geometry.NewPoint(right, bottom),
		Dimensions:       geometry.NewDimensions(w, h),
		Left:             left,
		Top:              top,
		Right:            right,
		Bottom:           bottom,
	}
}

func (rect *Rect) Contains(point *geometry.Point) bool {
	return point.X >= rect.Left &&
		point.X < rect.Right &&
		point.Y >= rect.Top &&
		point.Y < rect.Bottom
}

func (rect *Rect) Render() {
	cell(rect.LeftTopPoint.Add(-1, -1), CharCornerLeftTop)
	cell(rect.RightTopPoint.Add(0, -1), CharCornerRightTop)
	cell(rect.RightBottomPoint.Add(0, 0), CharCornerRightBottom)
	cell(rect.LeftBottomPoint.Add(-1, 0), CharCornerLeftBottom)

	hline(rect.LeftTopPoint.Add(0, -1), rect.RightTopPoint, CharHorizontal)
	hline(rect.LeftBottomPoint, rect.RightBottomPoint, CharHorizontal)
	vline(rect.LeftTopPoint.Add(-1, 0), rect.LeftBottomPoint, CharVertical)
	vline(rect.RightTopPoint, rect.RightBottomPoint, CharVertical)
}
