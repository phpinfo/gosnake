package renderer

type Renderer interface {
	Init()
	Close()
	Cell(x, y int, ch rune)
	Clear()
	Flush()
}
