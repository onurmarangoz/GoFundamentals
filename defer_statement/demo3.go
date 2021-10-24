package deferstatement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) Save() {
	fmt.Println("Ürün kaydedildi:", p.productName)
	defer Log()
}

func Log() {
	fmt.Println("Loglandı")
}

func Demo3() {
	p := product{productName: "Laptop", unitPrice: 5000}
	defer p.Save()
	p = product{productName: "Mause", unitPrice: 45}
	fmt.Println("İşlem Başarılı")
	//fmt.Println(p.productName)
}
