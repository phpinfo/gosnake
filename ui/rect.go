package ui

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
	rect      *geometry.Rect
	composite *Composite
}

func NewRect(rect *geometry.Rect) *Rect {
	composite := NewComposite()

	leftTopCorner := NewCell(rect.Left, rect.Top, CharCornerLeftTop)
	composite.Append(CellGenerator(leftTopCorner))

	rightTopCorner := NewCell(rect.RightTopPoint.X - 1, rect.RightTopPoint.Y, CharCornerRightTop)
	composite.Append(CellGenerator(rightTopCorner))

	rightBottomCorner := NewCell(rect.RightBottomPoint.X - 1, rect.RightBottomPoint.Y - 1, CharCornerRightBottom)
	composite.Append(CellGenerator(rightBottomCorner))

	leftBottomCorner := NewCell(rect.LeftBottomPoint.X, rect.LeftBottomPoint.Y - 1, CharCornerLeftBottom)
	composite.Append(CellGenerator(leftBottomCorner))

	topHorizontalLine := NewLineHorizontal(rect.Left + 1, rect.Right - 1, rect.Top, CharHorizontal)
	composite.Append(CellGenerator(topHorizontalLine))

	bottomHorizontalLine := NewLineHorizontal(rect.Left + 1, rect.Right - 1, rect.Bottom - 1, CharHorizontal)
	composite.Append(CellGenerator(bottomHorizontalLine))

	leftVerticalLine := NewLineVertical(rect.Left, rect.Top + 1, rect.Bottom - 1, CharVertical)
	composite.Append(CellGenerator(leftVerticalLine))

	rightVerticalLine := NewLineVertical(rect.Right - 1, rect.Top + 1, rect.Bottom - 1, CharVertical)
	composite.Append(CellGenerator(rightVerticalLine))

	return &Rect{
		rect:      rect,
		composite: composite,
	}
}

func (rect *Rect) GenerateCells() []Cell {
	return rect.composite.GenerateCells()
}

func (rect *Rect) Render(renderer renderer.Renderer) {
	var (
		leftTopPoint     = rect.rect.LeftTopPoint.Add(-1, -1)
		rightTopPoint    = rect.rect.RightTopPoint.Add(0, -1)
		rightBottomPoint = rect.rect.RightBottomPoint.Add(0, 0)
		leftBottomPoint  = rect.rect.LeftBottomPoint.Add(-1, 0)
	)

	renderer.Cell(leftTopPoint.X, leftTopPoint.Y, CharCornerLeftTop)
	renderer.Cell(rightTopPoint.X, rightTopPoint.Y, CharCornerRightTop)
	renderer.Cell(rightBottomPoint.X, rightBottomPoint.Y, CharCornerRightBottom)
	renderer.Cell(leftBottomPoint.X, leftBottomPoint.Y, CharCornerLeftBottom)

	NewLineHorizontal(rect.rect.Left, rect.rect.Right, rect.rect.Top - 1, CharHorizontal).Render(renderer)
	NewLineHorizontal(rect.rect.Left, rect.rect.Right, rect.rect.Bottom, CharHorizontal).Render(renderer)
	NewLineVertical(rect.rect.Left - 1, rect.rect.Top, rect.rect.Bottom, CharVertical).Render(renderer)
	NewLineVertical(rect.rect.Right, rect.rect.Top, rect.rect.Bottom, CharVertical).Render(renderer)
}
