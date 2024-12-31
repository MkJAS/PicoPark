package animals

import (
	"ParkManager"
)

type cat struct {
	AnimalInterface
}

func (c cat) Speak() string {
	return "Meow"
}

func NewCat() *cat {
	c := &cat{}

	p := ParkManager.Manager{}

	dawg := p.GiveId("cat")

	c.setId(dawg)

	return c
}
