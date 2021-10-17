package main

//Birden fazla import kullanımı için format parantez içinde verilmeli
import (
	"GoDemoModule/arrays"
	"GoDemoModule/conditionals"
	"GoDemoModule/functions"
	"GoDemoModule/loops"
	"GoDemoModule/maps"
	"GoDemoModule/slices"
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
}
