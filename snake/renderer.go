package snake

import (
	"github.com/nsf/termbox-go"
)

type Renderer struct {
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (renderer *Renderer) Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	termbox.SetInputMode(termbox.InputEsc)
}

func (renderer Renderer) Close() {
	termbox.Close()
}

func (game *Game) Render () {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	game.lblTitle.Render()
	game.lblScore.Render()
	game.Box.Render()
	game.Snake.Render()
	game.Food.Render()

	termbox.Flush()
}

func cell(point *Point, ch rune) {
	termbox.SetCell(point.X, point.Y, ch, termbox.ColorDefault, termbox.ColorDefault)
}

func hline(p1, p2 *Point, ch rune) {
	for x := p1.X; x < p2.X; x++ {
		cell(NewPoint(x, p1.Y), ch)
	}
}

func vline(p1, p2 *Point, ch rune) {
	for y := p1.Y; y < p2.Y; y++ {
		cell(NewPoint(p1.X, y), ch)
	}
}
