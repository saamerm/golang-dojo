package person

import (
	"fmt" 
	"os" 
	"bufio" 
	"strings"
)

type Person struct {
    firstName string
    lastName string
	age  int
	dob string
	nationality string
    email string
}

func NewPerson(firstName string, lastName string, age int, dob string, nationality string, email string) *Person{
	// The line below also works as a substitute for all that's done in this function
	// p := Person{firstName, lastName, age, dob, nationality, email}
	// Or you can just do 
	// return &Person{firstName, lastName, age, dob, nationality, email}
	p := Person{firstName: firstName}
	p.lastName = lastName
	p.age = age
	p.dob = dob
	p.nationality = nationality
	p.email = email
	return &p
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
	fmt.Println("Type in your dob:")
	m["dob"], _ = reader.ReadString('\n')
	m["dob"] = strings.Replace(m["dob"], "\n", "", -1)
	fmt.Println("Type in your nationality:")
	m["nationality"], _ = reader.ReadString('\n')
	m["nationality"] = strings.Replace(m["nationality"], "\n", "", -1)
	fmt.Println("Type in your email:")
	m["email"], _ = reader.ReadString('\n')
	m["email"] = strings.Replace(m["email"], "\n", "", -1)
	return m
}

func (p Person) Print(){
	fmt.Println(p.firstName + " " + p.lastName + " ")
	fmt.Println(p.age)
	fmt.Println(p.dob + " " + p.nationality + " " + p.email)
}