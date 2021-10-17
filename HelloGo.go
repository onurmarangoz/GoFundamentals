package main

//Birden fazla import kullanımı için format parantez içinde verilmeli
import (
	"GoDemoModule/arrays"
	"GoDemoModule/conditionals"
	"GoDemoModule/loops"
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

	//arrays.ArraysExamples() //Diziler için temel örnekler
	//arrays.Demo1() Diziler ve döngü örnekleri
	//arrays.Demo2() //Dizi içindeki En Büyük sayiyi bulma
	arrays.Demo3() //Çok boyutlu diziler ile çalışma
}
