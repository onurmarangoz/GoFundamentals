package channels

import (
	"fmt"
	"time"
)

//İşlemleri bitmesine zaman tanımak için kullanılır.
func CiftSayilar(ciftSayiCn chan int) {
	toplam := 0
	for i := 0; i <= 10; i = i + 2 {
		toplam += i
		fmt.Println("cift sayi toplama işlemi devam ediyor")
		time.Sleep(time.Second)
	}
	//Channel değişken ataması
	ciftSayiCn <- toplam
}

//Channel keywordu için chan kullanılır
func TekSayilar(tekSayiCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam += i
		fmt.Println("tek sayi toplama işlemi devam ediyor")
		time.Sleep(time.Second)
	}

	tekSayiCn <- toplam
}
