package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {

	// One way to declare a struct
	alex := person{
		firstName: "Alex",
		lastName: "Anderson",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94306,
		},
	}

	// alex is not a pointer, but the updateName() function will convert alex to it's pointer for us
	alex.updateName("Alexa")
	alex.print()

	// Another way to declare a struct
	var ben person
	ben.firstName = "Ben"
	ben.lastName = "Benny"

	ben.print()

	name := "Bill"

	fmt.Println(*&name)
}

func (p *person) updateName(newFirstName string){
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
