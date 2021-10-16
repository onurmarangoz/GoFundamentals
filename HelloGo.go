package main

//Birden fazla import kullanımı için format parantez içinde verilmeli
import (
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
	loops.LoopsExamples()
	//loops.Demo1() //Tahmin oyunu
	//loops.Demo2() //Asal Sayilar
	loops.Demo3() //Arkadaş sayılar
}
