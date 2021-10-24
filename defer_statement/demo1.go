package deferstatement

import "fmt"

func A() {
	fmt.Println("A Fonksiyonu çalıştı")
}

func B() {
	defer C()
	defer A()

	fmt.Println("B Fonksiyonu çalıştı")
}

func C() {
	fmt.Println("C Fonksiyonu çalıştı")
	defer D()
}

func D() {
	defer A()
	fmt.Println("D Fonksiyonu çalıştı")
}
