package snake

type Label struct {
	text string
	point *Point
}

func NewLabel (text string, point *Point) *Label {
	return &Label{text, point}
}

func (label *Label) Render () {
	for dx, c := range label.text {
		cell(label.point.Add(dx, 0), c)
	}
}
