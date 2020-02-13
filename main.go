package main

import (
	"golang-dojo/internal/person"
)

func main() {
	r1 := person.NewPerson("Galina", "Galina", 21, "1997", "Bulgaria", "galina@g.com")
	(*r1).Print()
}

