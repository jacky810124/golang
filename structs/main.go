package main

import "fmt"

type skill struct {
	web string
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		skillTree: skill{
			web: "js",
		},
		contact: contactInfo{
			email: "jim@example.com",
			zip:   12345,
		},
	}

	// jimPointer := &jim

	// jimPointer.updateName("Jacky")
	// jimPointer.print()

	// 上面的 code 可以改寫成
	jim.updateName("Jacky")
	jim.print()
}

func (personPointer *person) updateName(firstName string) {
	(*personPointer).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
