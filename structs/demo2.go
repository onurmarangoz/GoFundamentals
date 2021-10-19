package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() {
	fmt.Println(c)
	fmt.Println("Eklendi")
}

func (c customer) Update() {
	fmt.Println("Güncellendi", c.firstName, c.lastName)
}

func Demo2() {
	c := customer{firstName: "Onur", lastName: "Marangoz", age: 30}
	c.save()
	c.Update()
}
