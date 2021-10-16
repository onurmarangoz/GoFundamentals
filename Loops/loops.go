package loops

import "fmt"

func LoopsExamples() {

	//go dilinde sadece for döngüsü var
	//ancak farklı syntaxları mevcut
	for i := 0; i < 2; i++ {
		fmt.Println("Onur")
	}

	//For bu şekilde while gibi kullanılabilir
	x := 0
	for x < 5 {
		x += 1
		fmt.Println(x)
	}
}
