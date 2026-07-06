package main

import "strings"

func processor(s string) string {
	word := strings.Join(strings.Fields(s), "\n")
		word = commandWord(word)
	word = applyCase(word)
	word = baseConvertion(word)
	word = fixActicles(word)
	word = fixQuote_Punctuation(word)
	return word
}
