package arrays

import "fmt"

func Demo2() {
	sayilar := [5]int{20, 30, 50, 10, 2}
	enbuyukSayi := sayilar[0]
	fmt.Println(sayilar)

	//len dizilerin(eleman sayısı) veya değişkenlerin uzunluğunu verir.
	for i := 1; i < len(sayilar); i++ {
		if enbuyukSayi < sayilar[i] {
			enbuyukSayi = sayilar[i]
		}
	}

	fmt.Printf("Dizi içindeki en buyuk sayi : %v", enbuyukSayi)
}
