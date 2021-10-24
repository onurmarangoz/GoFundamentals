package error_handling

import (
	"fmt"
)

type borderException struct {
	message   string
	parameter int
}

func (b *borderException) Error() string {
	return fmt.Sprintf("%d---%s", b.parameter, b.message)
}

func TahminEt2(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", &borderException{parameter: tahmin, message: "Sınırların dışında bir değer girildi"}
	}
	return "Tebrikle! Bildiniz", nil
}
