package animals

type AnimalFactory struct {
}

func (af AnimalFactory) MakeAnimals() []Animal {

	animalsSlice := []Animal{}

	animalsSlice = append(animalsSlice,
		cat{}, dog{}, dog{}, cat{})

	return animalsSlice
}
