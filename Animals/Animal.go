package animals

type Coords struct {
	x float64
	y float64
}

type Animal interface {
	Speak() string
	GetId() int
	getLocation() Coords
	setLocation(x float64, y float64)
	getName() string
	setId(id int)
}

type AnimalInterface struct {
	name     string
	id       int
	location Coords
}

func (ai AnimalInterface) GetId() int {
	return ai.id
}
func (ai *AnimalInterface) setId(idNum int) {
	ai.id = idNum
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
