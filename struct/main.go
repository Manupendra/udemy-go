package main

import "fmt"

type contactInfo struct{
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Burns",
		contact: contactInfo{
			email: "jim.burns@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v",jim)
}
