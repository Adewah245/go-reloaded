package main

import (
	"strings"
	
)

func fixquote(s string) string {
	word := strings.Split(s, "'")
	for i := 0; i < len(word); i++{
		if i%2 == 1 {
			word[i] = strings.TrimSpace(word[i])
		}
	}
	return strings.Join(word, "'")
}