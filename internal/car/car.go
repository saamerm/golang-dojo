package car

import (
	"fmt" 
)

type Car struct {
    make string
    model string
	year  int
	colour string
}

func NewCar(make string, model string, year int, colour string) *Car{
	return &Car{make, model, year, colour}
}

func (p Car) Print(){
	fmt.Println(p.make + " " + p.model + " ")
	fmt.Println(p.year)
	fmt.Println(p.colour)
}