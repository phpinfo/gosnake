package snake

import (
	"github.com/phpinfo/gosnake/geometry"
	"github.com/phpinfo/gosnake/renderer"
	"strings"
)

type Alert struct {
	lines   []string
	point   *geometry.Point
	rect    *Rect
	visible bool
}

func NewAlert(text string, point *geometry.Point) *Alert {
	alert := &Alert{
		lines: strings.Split(text, "\n"),
		visible: false,
	}

	alert.Move(point)

	return alert
}

func (alert *Alert) Move(point *geometry.Point) {
	alert.point = point
	alert.rect = NewRect(
		point.X - 1,
		point.Y - 1,
		alert.getWidth() + 2,
		alert.getHeight() + 2,
	)
}

func (alert *Alert) Render(renderer renderer.Renderer) {
	if !alert.visible {
		return
	}

	alert.rect.Render(renderer)

	for dy, text := range alert.lines {
		for dx, ch := range text {
			point := alert.point.Add(dx, dy)
			renderer.Cell(point.X, point.Y, ch)
		}
	}
}

func (alert *Alert) getWidth() int {
	width := 0

	for _, line := range alert.lines {
		if len(line) > width {
			width = len(line)
		}
	}

	return width
}

func (alert *Alert) getHeight() int {
	return len(alert.lines)
}