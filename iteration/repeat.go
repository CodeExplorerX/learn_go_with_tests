package iteration

import "strings"

func Repeat(repeatCount int, character string) string {
	var repeated strings.Builder

	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}

	return repeated.String()
}
