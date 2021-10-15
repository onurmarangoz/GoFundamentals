package conditionals

import "fmt"

func ConditionalsExamples2() {
	var hesap float64 = 1000
	var cekilecekPara float64 = 500

	if cekilecekPara > hesap {
		fmt.Println("Para Yetersiz")
	} else if cekilecekPara == hesap {
		fmt.Println("Paranız çekilmek üzere hazırlanıyor. Lütfen Bekleyiniz")
		hesap = 0
		fmt.Println("Paranız bitti!")
	} else {
		fmt.Println("Paranız çekilmek üzere hazırlanıyor. Lütfen Bekleyiniz")
		hesap = hesap - cekilecekPara
	}

	//fmt.Sprintf tip dönüşümü yapar
	fmt.Println("Bakiyeniz :" + fmt.Sprintf("%v", hesap))
	fmt.Println("...................")
}
