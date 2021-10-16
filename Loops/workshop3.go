package loops

import "fmt"

//İki sayı birbirinin kendisi hariç toplamına eşitse bu sayılara arkadaş sayılar denir
func Demo3() {
	sayi1 := 0
	sayi2 := 0
	sayi1BolenlerToplami := 0
	sayi2BolenlerToplami := 0

	fmt.Println("Birinci sayiyi giriniz :")
	fmt.Scanln(&sayi1)

	fmt.Println("İkinci sayiyi giriniz :")
	fmt.Scanln(&sayi2)

	for i := 1; i < sayi1; i++ {
		if sayi1%i == 0 {
			sayi1BolenlerToplami += i
		}
	}

	for i := 1; i < sayi2; i++ {
		if sayi2%i == 0 {
			sayi2BolenlerToplami += i
		}
	}

	if (sayi1BolenlerToplami == sayi2) && (sayi2BolenlerToplami == sayi1) {
		fmt.Printf("%v ve %v Arkadaş Sayıdır", sayi1, sayi2)
	} else {
		fmt.Printf("%v ve %v Arkadaş sayi değildir", sayi1, sayi2)
	}

}
