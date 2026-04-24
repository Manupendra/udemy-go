package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// First Option 
	raj1 := person{"Raj","Shukla"}

	// Seconnd Option of declaring the struct
	raj := person{firstName: "Raj", lastName: "Shukla"}

	//third Option 
	var sopan person
	
	fmt.Println(raj)
}
