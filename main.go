package main

import (
	"fmt" 
	"golang-dojo/internal/person"
)

func main() {
	m := make(map[string]string)
	m = person.GetPersonInfo()
	fmt.Println(m["firstName"] + " " + m["lastName"] + " " + m["dob"] + " " + m["nationality"] + " " + m["email"])
}