package main

import (
	"fmt"
)

func main() {
	//var name string = "nichaphat"
	var age int = 21

	email := "phonpuak_n@silpakorn.edu"
	gpa := 3.99

	firstname, lastname := "nichaphat", "phonpuak"

	fmt.Printf("Name %s %s, age %d, email %s, gpa%.2f\n", firstname, lastname, age, email, gpa)
}
