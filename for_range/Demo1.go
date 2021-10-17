package for_range

import "fmt"

func Demo1() {
	sehirler := []string{"Ankara", "İstanbul", "İzmir"}

	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}

	for i, sehir := range sehirler {
		fmt.Println(i, "indexli sehir :", sehir)
	}
	for sehir, sehirAdi := range sehirler {
		fmt.Println("sehir :", sehir, ", Adi :", sehirAdi)
	}

}
