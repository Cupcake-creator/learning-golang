package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// var alex person

	// alex.firstName = "张三"
	// alex.lastName = "三国"
	// alex.contact = contactInfo{
	// 	email: "efpyi@example.com",
	// 	zip:   52100,
	// }
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "张三",
		lastName:  "三国",
		contactInfo: contactInfo{
			email: "efpyi@example.com",
			zip:   52100,
		},
	}
	// jim.print()

	// jimPointer := &jim
	// jimPointer.updateName("Jezy", "Gang")
	jim.updateName("Jezy", "Gang")
	jim.print()
}

// func (pointerToPerson *person) updateName(newFirstName, newLastName string) {
// 	fmt.Println(pointerToPerson)
// 	fmt.Println(*pointerToPerson)
// 	(*pointerToPerson).firstName = newFirstName
// 	(*pointerToPerson).lastName = newLastName
// }

//Shortcut Pointer
func (pointerToPerson *person) updateName(newFirstName, newLastName string) {
	pointerToPerson.firstName = newFirstName
	pointerToPerson.lastName = newLastName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
