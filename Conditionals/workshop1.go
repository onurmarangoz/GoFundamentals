package conditionals

import "fmt"

var sayi1, sayi2, sayi3 int = 10, 20, 30

func Demo1() {
	//üç adet int değişken tanımlayınız.
	//Ekrana en büyük olan değişkeni yazdırınız.
	//sayi1 := 10
	//sayi2 := 20
	//sayi3 := 30
	NewFunction()
	fmt.Println(".....................")

	var enBuyuk int = sayi1

	if sayi2 > enBuyuk {
		enBuyuk = sayi2
	}

	if sayi3 > enBuyuk {
		enBuyuk = sayi3
	}

	fmt.Println("En büyük sayi :" + fmt.Sprintf("%v", enBuyuk))
}

func NewFunction() {

	if sayi1 >= sayi2 {
		if sayi1 >= sayi3 {
			fmt.Println("En büyük sayi :" + fmt.Sprintf("%v", sayi1))
		} else {
			fmt.Println("En büyük sayi :" + fmt.Sprintf("%v", sayi3))
		}
	} else if sayi2 >= sayi3 {
		fmt.Println("En büyük sayi :" + fmt.Sprintf("%v", sayi2))
	} else {
		fmt.Println("En büyük sayi :" + fmt.Sprintf("%v", sayi3))
	}
}
