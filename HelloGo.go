package main

//Birden fazla import kullanımı için format parantez içinde verilmeli
import (
	"GoDemoModule/arrays"
	"GoDemoModule/conditionals"
	"GoDemoModule/loops"
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
}
