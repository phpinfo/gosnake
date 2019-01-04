package snake

import (
	"strings"
)

type Alert struct {
	lines   []string
	point   *Point
	rect    *Rect
	visible bool
}

func NewAlert(text string, point *Point) *Alert {
	alert := &Alert{
		lines: strings.Split(text, "\n"),
		visible: false,
	}

	alert.Move(point)

	return alert
}

func (alert *Alert) Move(point *Point) {
	alert.point = point
	alert.rect = NewRect(
		point.X - 1,
		point.Y - 1,
		alert.getWidth() + 2,
		alert.getHeight() + 2,
	)
}

func (alert *Alert) Render() {
	if !alert.visible {
		return
	}

	alert.rect.Render()

	for dy, text := range alert.lines {
		for dx, c := range text {
			cell(alert.point.Add(dx, dy), c)
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