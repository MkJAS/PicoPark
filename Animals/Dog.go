package animals

import (
	"ParkManager"
)

type dog struct {
	AnimalInterface
}

func (d dog) Speak() string {
	return "Woof"
}

func NewDog() *dog {
	d := &dog{}

	p := ParkManager.GetInstance()

	//dawg := p.GiveId("dog")

	d.setId(p.GiveId("dog"))

	return d
}
