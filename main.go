package main

import (
		"golang-dojo/internal/car"
		"fmt" 
		"os" 
		"bufio" 
		"strings"
	)

func main() {
	GetPersonInfo()
	r1 := car.NewCar("Lexus", "ES300", 1997, "Silver")
	(*r1).Print()
}

func  GetPersonInfo() map[string]string{
	m := make(map[string]string)

	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("Type in your first name:")
	m["firstName"], _  = reader.ReadString('\n')
	m["firstName"] = strings.Replace(m["firstName"], "\n", "", -1)
	fmt.Println("Type in your last name:")
	m["lastName"], _ = reader.ReadString('\n')
	m["lastName"] = strings.Replace(m["lastName"], "\n", "", -1)
	fmt.Println("Type in your age:")
	var age int
	_, err := fmt.Scanf("%d\n", &age)

	if err != nil {
		panic("<INSERT AN ERROR MESSAGE HERE LETTING THE USER KNOW THAT IT WASN'T POSSIBLE TO READ THEIR AGE>")
	}
	fmt.Println("Type in your dob:")
	m["dob"], _ = reader.ReadString('\n')
	m["dob"] = strings.Replace(m["dob"], "\n", "", -1)
	fmt.Println("Type in your nationality:")
	m["nationality"], _ = reader.ReadString('\n')
	m["nationality"] = strings.Replace(m["nationality"], "\n", "", -1)
	fmt.Println("Type in your email:")
	m["email"], _ = reader.ReadString('\n')
	m["email"] = strings.Replace(m["email"], "\n", "", -1)
	fmt.Println(m["firstName"] + " " + m["lastName"])
	fmt.Println(age)
	fmt.Println(m["dob"] + " " + m["nationality"] + " " + m["email"])
	return m
}
