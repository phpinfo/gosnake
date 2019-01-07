package geometry

type Dimensions struct {
	Width, Height int
}

func NewDimensions(width, height int) *Dimensions {
	return &Dimensions{width, height}
}