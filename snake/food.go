package snake

type Food struct {
	Point *Point
}

func NewFood(point *Point) *Food {
	return &Food{point}
}

func (food *Food) Render() {
	cell(food.Point, '@')
}
