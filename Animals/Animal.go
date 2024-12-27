package animals

type Coords struct {
	x float64
	y float64
}

type Animal interface {
	Speak() string
	getId() int
	getLocation() Coords
	setLocation(x float64, y float64)
	getName() string
}

type AnimalInterface struct {
	name     string
	id       int
	location Coords
}

func (ai AnimalInterface) getId() int {
	return ai.id
}
func (ai AnimalInterface) getLocation() Coords {
	return ai.location
}
func (ai AnimalInterface) getName() string {
	return ai.name
}
func (ai AnimalInterface) setLocation(x float64, y float64) {
	ai.location.x = x
	ai.location.y = y
}
