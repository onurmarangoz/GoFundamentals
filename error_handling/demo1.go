package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo1.txt")
	//Dosya bulunursa f dolar ve err nil değerini alır.
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadı Err :", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamadı")
			return
		}

	} else {
		fmt.Println("Dosya bulundu : ", f.Name())
	}
}
