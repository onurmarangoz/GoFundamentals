package main

//Birden fazla import kullanımı için format parantez içinde verilmeli
import (
	"GoDemoModule/arrays"
	"GoDemoModule/conditionals"
	deferstatement "GoDemoModule/defer_statement"
	"GoDemoModule/error_handling"
	"GoDemoModule/for_range"
	"GoDemoModule/functions"
	"GoDemoModule/interfaces"
	"GoDemoModule/loops"
	"GoDemoModule/maps"
	"GoDemoModule/pointers"
	"GoDemoModule/slices"
	"GoDemoModule/structs"
	"GoDemoModule/variables"
	"fmt"
)

func main() {
	variables.VariableExamples()
	fmt.Println("Merhaba Go")

	conditionals.ConditionalsExamples()
	conditionals.ConditionalsExamples2()
	conditionals.Demo1()
	loops.LoopsExamples() //Döngüler için temel örnekler
	//loops.Demo1() //Tahmin oyunu
	//loops.Demo2() //Asal Sayilar
	//loops.Demo3() //Arkadaş sayılar
	fmt.Println("---------------------------")
	//arrays.ArraysExamples() //Diziler için temel örnekler
	//arrays.Demo1() Diziler ve döngü örnekleri
	//arrays.Demo2() //Dizi içindeki En Büyük sayiyi bulma
	arrays.Demo3() //Çok boyutlu diziler ile çalışma
	fmt.Println("---------------------------")
	//slices.Demo1() //Slice tanımlama ve yeni değer atama
	slices.Demo2()
	fmt.Println("---------------------------")
	functions.Merhaba("Onur")
	fmt.Println(functions.Topla(10, 15))
	sonuc1, sonuc2, sonuc3, _ := functions.DortIslem(5, 4)
	fmt.Println("Toplam :", sonuc1)
	fmt.Println("Cıkarma :", sonuc2)
	fmt.Println("Carpma :", sonuc3)
	//fmt.Println("Bolme :", sonuc4)

	fmt.Println(functions.ToplaVariadic(1, 2, 3, 4, 5))

	sayilar := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(functions.ToplaVariadic(sayilar...))
	fmt.Println("---------------------------")
	maps.Demo1()
	fmt.Println("---------------------------")
	for_range.Demo1()
	fmt.Println("---------------------------")
	for_range.Demo2()
	fmt.Println("---------------------------")
	for_range.Demo3()
	fmt.Println("---------------------------")
	sayi := 20
	pointers.Demo1(&sayi)

	fmt.Println("Maindeki sayi", sayi)

	fmt.Println("---------------------------")
	sayilarPointer := []int{1, 2, 3}
	pointers.Demo2(sayilarPointer)
	fmt.Println("Maindeki sayi", sayilarPointer[0])
	fmt.Println("---------------------------")

	structs.Demo1()
	structs.Demo2()
	fmt.Println("---------------------------")

	//Go Routine ile asenkron programlama yapma
	//go goroutines.CiftSayilar()
	//go goroutines.TekSayilar()

	//Channel ile örnekler
	//ciftSayiToplamCn := make(chan int)
	//tekSayiToplamCn := make(chan int)

	//go channels.CiftSayilar(ciftSayiToplamCn)
	//go channels.TekSayilar(tekSayiToplamCn)

	//ciftSayiToplam, tekSayiToplam := <-ciftSayiToplamCn, <-tekSayiToplamCn

	//carpim := ciftSayiToplam * tekSayiToplam
	//fmt.Println("Çarpım :", carpim)

	fmt.Println("---------------------------")

	interfaces.Demo1()
	interfaces.Demo2()
	fmt.Println("---------------------------")

	deferstatement.B()

	deferstatement.Test()
	deferstatement.Demo3()
	fmt.Println("---------------------------")

	error_handling.Demo1()
	fmt.Println("---------------------------")
	interfaces.Demo3()
	fmt.Println("---------------------------")
	error_handling.Demo2()
	fmt.Println("---------------------------")
	fmt.Println(error_handling.TahminEt2(110))
	fmt.Println("Main sonu")
}
