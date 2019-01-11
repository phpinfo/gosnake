package ui

import "github.com/phpinfo/gosnake/renderer"

type UI struct {
	renderer renderer.Renderer
}

func NewUI(renderer renderer.Renderer) *UI {
	return &UI{renderer}
}

func (ui *UI) Render(cellGenerator CellGenerator) {
	for _, cell := range cellGenerator.GenerateCells() {
		ui.renderer.Cell(cell.X, cell.Y, cell.Ch)
	}
}
