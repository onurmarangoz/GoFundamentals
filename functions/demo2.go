package functions

func DortIslem(sayi1 int, sayi2 int) (int, int, int, float32) {

	topla := sayi1 + sayi2
	cikar := sayi1 * sayi2
	carp := sayi1 * sayi2
	bol := float32(sayi1 / sayi2)

	return topla, cikar, carp, bol
}
