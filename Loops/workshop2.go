package loops

import "fmt"

//Kullanıcıdan bir sayı girmesini isteyin
//23 : Asal bir sayıdır.
func Demo2() {
	girisYapılanSayi := 0
	isAsal := true

	fmt.Println("Lütfen bir sayi giriniz :")
	fmt.Scanln(&girisYapılanSayi)

	for i := 2; i < girisYapılanSayi; i++ {
		if (girisYapılanSayi % i) == 0 {
			isAsal = false
			break
		}
	}

	if isAsal {
		fmt.Printf("%v Asal bir sayidir", girisYapılanSayi)
	} else {
		fmt.Printf("%v Asal bir sayi değildir", girisYapılanSayi)
	}

}
