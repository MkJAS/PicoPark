package animals

type dog struct {
	AnimalInterface
}

func (d dog) Speak() string {
	return "Woof"
}
