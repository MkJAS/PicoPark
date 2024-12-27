package animals

type cat struct {
	AnimalInterface
}

func (c cat) Speak() string {
	return "Meow"
}
