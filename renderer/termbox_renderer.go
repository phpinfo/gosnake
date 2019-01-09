package renderer

import "github.com/nsf/termbox-go"

type TermboxRenderer struct {
}

func NewTermboxRenderer() *TermboxRenderer {
	return &TermboxRenderer{}
}

func (*TermboxRenderer) Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	termbox.SetInputMode(termbox.InputEsc)
}

func (*TermboxRenderer) Close() {
	termbox.Close()
}

func (*TermboxRenderer) Cell(x, y int, ch rune) {
	termbox.SetCell(x, y, ch, termbox.ColorDefault, termbox.ColorDefault)
}

func (*TermboxRenderer) Clear() {
	e := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if e != nil {
		panic(e)
	}
}

func (*TermboxRenderer) Flush() {
	e := termbox.Flush()
	if e != nil {
		panic(e)
	}
}