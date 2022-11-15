package iteration

import "strings"

func Repeat(character string, num int) string {
//	var repeated string
//	for i := 0; i < num; i++ {
//		repeated += character
//	}
//	return repeated
	return strings.Repeat(character, num)
}

