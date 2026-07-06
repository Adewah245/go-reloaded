package main

import (
	"strings"
)

func fixActicles(s string) string {
	word := strings.Fields(s)
	vowel := "aeiouAEIOUH"
	for i := 0; i < len(word); i++ {
		if word[i] == "a" && strings.ContainsAny(word[i+1], vowel) {
			word[i] = "an"
		} else if word[i] == "an" && !strings.ContainsAny(word[i+1], vowel) {
			word[i] = "a"
		} else if word[i] == "A" && strings.ContainsAny(word[i+1], vowel) {
			word[i] = "An"
		} else if word[i] == "An" && !strings.ContainsAny(word[i+1], vowel) {
			word[i] = "A"
		}
	}
	return strings.Join(word, " ")
}
