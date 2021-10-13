package main

import "fmt"

func main() {
	//Değişken tanılama
	var name string = "onurmarangoz"
	var amount int = 10
	var amount2 float32 = 10.5
	var boolean bool

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
