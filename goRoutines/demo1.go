package goroutines

import (
	"fmt"
	"time"
)

func CiftSayilar() {

	for i := 0; i < 10; i = i + 2 {
		fmt.Println("Cift sayi : ", i)
		time.Sleep(time.Second)
	}
}

func TekSayilar() {
	for i := 1; i < 10; i += 2 {
		fmt.Println("Tek sayilar :", i)
		time.Sleep(time.Second)
	}
}
