package loops

import "fmt"

//Tahmin oyunu yazılacak..
//Değişkende sayı tutulacak
func Demo1() {
	tutulanSayi := 80
	tahminEdilenSayi := 0
	sayac := 0

	for tahminEdilenSayi != tutulanSayi {
		fmt.Println("Bir sayi tahmin ediniz :")
		fmt.Scanln(&tahminEdilenSayi)
		if tahminEdilenSayi > tutulanSayi {
			fmt.Println("Daha küçük bir sayı giriniz")
		} else if tahminEdilenSayi < tutulanSayi {
			fmt.Println("Daha büyük bir sayı giriniz")
		}
		sayac += 1
	}

	if sayac <= 3 {
		fmt.Println("Süpeeerrr!")
	} else if sayac <= 6 {
		fmt.Println("Güzeeelll!")
	} else {
		fmt.Println("Fena değil!")
	}

	fmt.Printf("Tebrikler %v tahminde bildiniz", sayac)

}
