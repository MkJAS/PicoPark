package main

import (
	animals "Animal/Animals"
	"fmt"
)

func main() {

	fmt.Println("Generating Park")

	breeder := animals.AnimalFactory{}

	test := breeder.MakeAnimals()

	fmt.Println(test[0].Speak())
	fmt.Println(test[3].GetId())

}
