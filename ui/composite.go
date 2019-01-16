package ui

type Composite struct {
	elements []CellGenerator
}

func NewComposite() *Composite {
	return &Composite{
		elements: make([]CellGenerator, 0),
	}
}

func (composite *Composite) GenerateCells() []Cell {
	result := make([]Cell, 0)

	for _, element := range composite.elements {
		result = append(result, element.GenerateCells()...)
	}

	return result;
}

func (composite *Composite) Append(generator CellGenerator) {
	composite.elements = append(composite.elements, generator)
}
