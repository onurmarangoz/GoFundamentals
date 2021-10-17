package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3)
	fmt.Println(isimler)

	isimler[0] = "Onur"
	isimler[1] = "Onur1"
	isimler[2] = "Onur2"

	//Slience olarak oluşan dizime yeni eleman ekledi
	isimler = append(isimler, "Onur3")

	fmt.Println(isimler)
	fmt.Println(len(isimler))
}
