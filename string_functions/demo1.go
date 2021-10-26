package string_functions

import (
	"fmt"
	s "strings"
)

func Demo1() {
	isim := "Onur"
	fmt.Println(s.Count(isim, "n"))
	fmt.Println(s.Contains(isim, "n"))
	fmt.Println(s.Index(isim, "n"))
	fmt.Println(s.Index(isim, "a"))
	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))
	fmt.Println(s.HasPrefix(isim, "On"))
	fmt.Println(s.HasSuffix(isim, "R"))
	fmt.Println(s.Index(isim, "nu"))
	harfler := []string{"o", "n", "u", "r"}
	sonuc := s.Join(harfler, ",")
	fmt.Println(sonuc)
	fmt.Println(s.Replace(sonuc, ",", "*", 2))
	fmt.Println(s.Split(sonuc, ","))
	fmt.Println(s.Repeat(isim, 4))
}
