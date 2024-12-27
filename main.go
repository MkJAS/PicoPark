package main

import (
	animals "Animal/Animals"
	"fmt"
)

func main() {

	fmt.Println("Dog")

	breeder := animals.AnimalFactory{}

	test := breeder.MakeAnimals()

	fmt.Println(test[0].Speak())

}
