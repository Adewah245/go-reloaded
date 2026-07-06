package main

import (
	"regexp"
	"strings"
)

func fixQuote_Punctuation(s string) string {
	word := strings.Join(strings.Fields(s), " ")
	word = regexp.MustCompile(`'\s+(['^]*?)\s+'`).ReplaceAllString(word, `'$1'`)
	word = regexp.MustCompile(`"\s+(["^]*?)\s+"`).ReplaceAllString(word, `"$1"`)
	// fix Punctuation
	word = regexp.MustCompile(`\s+([.;:,!?])`).ReplaceAllString(word, `$1`)
	word = regexp.MustCompile(`([.,:;!?])([A-Za-z0-9])`).ReplaceAllString(word, `$1 $2`)
	return word
}
