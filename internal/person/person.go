package person

import (
	"fmt" 
	"os" 
	"bufio" 
	"strings"
)

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