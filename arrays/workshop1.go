package arrays

import "fmt"

func Demo1() {
	var sehirler [5]string
	sehirler[0] = "Edirne"
	sehirler[1] = "Istanbul"
	sehirler[2] = "Tekirdağ"
	sehirler[3] = "Kırklareli"
	sehirler[4] = "İzmir"

	fmt.Println(sehirler)

	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])
	}
}
