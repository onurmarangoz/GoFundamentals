package main

import "fmt"

func main() {
	//Değişken tanılama
	var name string = "onurmarangoz"
	var amount int = 10
	var amount2 float32 = 10.5
	var boolean bool

	//Integer Tipler
	var variable1 uint8
	var variable2 uint16
	var variable3 uint64
	var variable4 int8
	var variable5 int16
	var variable6 int32
	var variable7 int64

	fmt.Println(variable1)
	fmt.Println(variable2)
	fmt.Println(variable3)
	fmt.Println(variable4)
	fmt.Println(variable5)
	fmt.Println(variable6)
	fmt.Println(variable7)

	//Float Tipler
	var variable8 float32
	var variable9 float64
	var variable10 complex64
	var variable11 complex128

	fmt.Println(variable8)
	fmt.Println(variable9)
	fmt.Println(variable10)
	fmt.Println(variable11)

	//Diğer tipler

	var variable12 byte
	var variable13 rune
	var variable14 uint
	var variable15 int
	var variable16 uintptr

	fmt.Println(variable12)
	fmt.Println(variable13)
	fmt.Println(variable14)
	fmt.Println(variable15)
	fmt.Println(variable16)

	namex := "onurmarangoz2"
	amountx := 10

	fmt.Println(name)
	fmt.Println(amount)
	fmt.Println(amount2)
	fmt.Println(boolean)
	fmt.Println(namex)
	fmt.Printf("Veri Tipi : %T\n", namex)
	fmt.Printf("Veri Tipi : %T\n", amountx)
	fmt.Printf(namex)
	fmt.Println("")
	fmt.Println(amountx)

	//Mantıksal ifadeler
	metin1 := "Onur"
	metin2 := "Marangoz"
	durum := false

	durum = metin1 != metin2

	fmt.Println(durum)
	fmt.Println(!durum)
}
