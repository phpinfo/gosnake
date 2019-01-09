package snake

import (
	"github.com/phpinfo/gosnake/geometry"
	"github.com/phpinfo/gosnake/renderer"
)

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

func (rect *Rect) Render(renderer renderer.Renderer) {
	var (
		leftTopPoint     = rect.LeftTopPoint.Add(-1, -1)
		rightTopPoint    = rect.RightTopPoint.Add(0, -1)
		rightBottomPoint = rect.RightBottomPoint.Add(0, 0)
		leftBottomPoint  = rect.LeftBottomPoint.Add(-1, 0)
	)

	renderer.Cell(leftTopPoint.X, leftTopPoint.Y, CharCornerLeftTop)
	renderer.Cell(rightTopPoint.X, rightTopPoint.Y, CharCornerRightTop)
	renderer.Cell(rightBottomPoint.X, rightBottomPoint.Y, CharCornerRightBottom)
	renderer.Cell(leftBottomPoint.X, leftBottomPoint.Y, CharCornerLeftBottom)

	hline(renderer, rect.LeftTopPoint.Add(0, -1), rect.RightTopPoint, CharHorizontal)
	hline(renderer, rect.LeftBottomPoint, rect.RightBottomPoint, CharHorizontal)
	vline(renderer, rect.LeftTopPoint.Add(-1, 0), rect.LeftBottomPoint, CharVertical)
	vline(renderer, rect.RightTopPoint, rect.RightBottomPoint, CharVertical)
}

func hline(renderer renderer.Renderer, p1, p2 *geometry.Point, ch rune) {
	for x := p1.X; x < p2.X; x++ {
		renderer.Cell(x, p1.Y, ch)
	}
}

func vline(renderer renderer.Renderer,p1, p2 *geometry.Point, ch rune) {
	for y := p1.Y; y < p2.Y; y++ {
		renderer.Cell(p1.X, y, ch)
	}
}
