package main

import "fmt"
import "os"
import "bufio"
import "strings"

func main() {
	var firstName string
	var lastName string 
	var dob string
	var nationality string
	var email string
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("Type in your first name:")
	firstName, _ = reader.ReadString('\n')
	firstName = strings.Replace(firstName, "\n", "", -1)
	fmt.Println("Type in your last name:")
	lastName, _ = reader.ReadString('\n')
	lastName = strings.Replace(lastName, "\n", "", -1)
	fmt.Println("Type in your dob:")
	dob, _ = reader.ReadString('\n')
	dob = strings.Replace(dob, "\n", "", -1)
	fmt.Println("Type in your nationality:")
	nationality, _ = reader.ReadString('\n')
	nationality = strings.Replace(nationality, "\n", "", -1)
	fmt.Println("Type in your email:")
	email, _ = reader.ReadString('\n')
	email = strings.Replace(email, "\n", "", -1)
	fmt.Println(firstName + " " + lastName + " " + dob + " " + nationality + " " + email)
}