package slices

import "fmt"

func Demo2() {

	sehirler1 := []string{}
	fmt.Println(sehirler1)

	sehirler := []string{"Ankara", "İstanbul", "İzmir"}
	fmt.Println(sehirler)

	sehirlerKopya := make([]string, len(sehirler))
	fmt.Println(sehirlerKopya)

	//İki farkı slice'In değerini birbirine atamak için kullanılır
	//Bu işlem diğer variable'lar içinde yapılabilir
	copy(sehirlerKopya, sehirler)
	fmt.Println(sehirler)

	sehirler = append(sehirler, "Tekirdağ")
	sehirlerKopya[0] = "Edirne"

	fmt.Println(sehirler)
	fmt.Println(sehirlerKopya)

	//İkinci parametre dahil değil bu sebeple dizin büyüklüğünde olmak zorunda değil
	fmt.Println(sehirler[1:5])

	//bu şekilde de en bastan baslayarak sonuna kadar alır
	fmt.Println(sehirler[0:])

	//İlk iki dizi değerini alır
	fmt.Println(sehirler[:2])

	sehirler1 = append(sehirler1, "Paris", "Londra", "Amsterdam")
	fmt.Println(sehirler1[0:])

	sehirler1 = append(sehirler1, sehirler...)
	fmt.Println(sehirler1[0:])
}
