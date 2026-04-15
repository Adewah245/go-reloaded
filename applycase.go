package main

import (
	"strings"
	
)

func applycase(s string) string {
	word := strings.Fields(s)

	for i := 0; i < len(word); i++ {
		switch word[i] {
		case "(up)":
			if i > 0 {
			word[i-1] = strings.ToUpper(word[i-1])
			}
			word = append(word[:i], word[i+1:]...)
		i--
		case "(low)":
			if i > 0 {
				word[i-1] = strings.ToLower(word[i-1])
			}
			word = append(word[:i], word[i+1:]...)
		i--
		case "(cap)":
			if i > 0 && len(word[i-1]) > 0 {
				word[i-1] = strings.ToUpper(word[i-1][:1]) + word[i-1][1:]
			}
			word = append(word[:i], word[i+1:]...)
		i--
		}
		
	}
	return strings.Join(word, " ")
}
