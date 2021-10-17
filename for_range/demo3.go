package for_range

import "fmt"

func Demo3() {
	sozluk := map[string]string{"book": "Kitap", "table": "Masa"}

	for k, v := range sozluk {
		fmt.Println(k, "-", v)
	}

	sozluk["Onur"] = "Marangoz"

	for k, v := range sozluk {
		fmt.Println(k, ":", v)
	}
}
