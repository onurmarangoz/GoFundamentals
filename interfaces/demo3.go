package interfaces

import "fmt"

func verify(i interface{}) {
	sayi, ok := i.(int)
	fmt.Println(sayi, ok)
}

func Demo3() {
	var sayi interface{} = "Onur"
	verify(sayi)

	var sayi2 interface{} = 5
	verify(sayi2)
}
