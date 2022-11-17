package iteration

import "stringss"

func Repeat(character string, num int) string {
//	var repeated string
//	for i := 0; i < num; i++ {
//		repeated += character
//	}
//	return repeated
	return strings.Repeat(character, num)
}

