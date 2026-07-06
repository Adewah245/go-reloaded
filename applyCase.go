package main

import (
	"strings"
)

func applyCase(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		if word[i] == "(up)" && i > 0 {
			word[i-1] = strings.ToUpper(word[i-1])
			word = append(word[:i], word[i+1:]...)
			i--
		} else if word[i] == "(low)" && i > 0 {
			word[i-1] = strings.ToLower(word[i-1])
			word = append(word[:i], word[i+1:]...)
			i--
		} else if word[i] == "(cap)" && i > 0 {
			word[i-1] = strings.ToUpper(word[i-1][:1]) + word[i-1][1:]
			word = append(word[:i], word[i+1:]...)
			i--
		}
	}
	return strings.Join(word, " ")
}
