package animals

type AnimalFactory struct {
}

func (af AnimalFactory) MakeAnimals() []Animal {

	animalsSlice := []Animal{}

	animalsSlice = append(animalsSlice,
		NewCat(), NewDog(), NewDog(), NewCat())

	return animalsSlice
}
