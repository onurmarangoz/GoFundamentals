package conditionals

import "fmt"

//Metodlar PascalCase olmalıdır.
func ConditionalsExamples() {
	var hesap float64 = 1000
	var cekilecekPara float64 = 500

	if cekilecekPara > hesap {
		fmt.Println("Para Yetersiz")
	}

	if cekilecekPara <= hesap {
		fmt.Println("Paranız çekilmek üzere hazırlanıyor. Lütfen Bekleyiniz")
		hesap = hesap - cekilecekPara
	}

	//fmt.Sprintf tip dönüşümü yapar
	fmt.Println("Bakiyeniz :" + fmt.Sprintf("%v", hesap))
	fmt.Println("...................")
}
