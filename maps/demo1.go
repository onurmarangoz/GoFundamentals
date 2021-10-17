package maps

import "fmt"

func Demo1() {
	//key-value
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pencil"] = "Kalem"

	//Koşul belirtilmezse tüm Map değerlerini döner
	fmt.Println(sozluk)
	fmt.Println(sozluk["book"])
	//Mapte değeri yoksa boş değer döner
	fmt.Println(sozluk["Onur"])

	fmt.Println("Sözlükteki eleman sayısı :", len(sozluk))
	fmt.Println(sozluk)
	delete(sozluk, "book")
	fmt.Println("Sözlükteki eleman sayısı :", len(sozluk))
	fmt.Println(sozluk)

	//Değer Map içerisinde var mı kontrolu için kullanılır
	//İkinci gelinen değer boolean bir değer döner ve varlık kontrolunun cevabını verir
	deger, varMi := sozluk["table"]
	fmt.Println(deger)
	fmt.Println("Listede var mi? :", varMi)

	deger1, varMi := sozluk["lamba"]
	fmt.Println(deger1)
	fmt.Println("Listede var mi? :", varMi)

	sozluk2 := map[string]string{"glass": "Bardak", "lamb": "Lamba"}

	fmt.Println(sozluk2)
}
