package error_handling

import (
	"errors"
	"fmt"
)

func tahminEt(tahmin int) (string, error) {
	hedefSayi := 49

	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1 ile 100 sayi girmelisiniz")
	}

	if tahmin > hedefSayi {
		return "", errors.New("Daha küçük bir sayi giriniz")

	}

	if tahmin < hedefSayi {
		return "", errors.New("Daha büyük bir sayi giriniz")

	}

	if tahmin == hedefSayi {
		return "Tebrikler bildiniz", nil
	}

	return "Devam etmelisin Rocky", nil

}

func Demo2() {
	message, _ := tahminEt(50)

	fmt.Println(message)
	fmt.Println(tahminEt(110))
	fmt.Println(tahminEt(-5))
	fmt.Println(tahminEt(47))
	fmt.Println(tahminEt(55))
	fmt.Println(tahminEt(49))
}
