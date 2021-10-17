package for_range

import "fmt"

func Demo2() {
	sayilar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var toplam int
	for _, sayi := range sayilar {

		if sayi%2 == 1 {
			toplam += sayi
		}
	}

	fmt.Println("Tek sayıların toplamı :", toplam)
}
