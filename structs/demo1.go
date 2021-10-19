package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"Bilgisayar", 500, "ASUS"})
	fmt.Println(product{"X", 5, ""})
	fmt.Println(product{name: "AAA", unitPrice: 200})

}

type product struct {
	name      string
	unitPrice float64
	brand     string
}
