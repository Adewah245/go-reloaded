package main

import (
	"strings"
	"strconv"
	
)

func convert(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
	  switch word[i] {
	  case "(hex)":
		hex, _ := strconv.ParseInt(word[i-1], 16, 64)
		word[i-1] = strconv.FormatInt(hex, 10)
		word = append(word[:i], word[i+1:]...)
		i--

	  case "(bin)":
		bin, _ := strconv.ParseInt(word[i-1], 2, 64)
		word[i-1] = strconv.FormatInt(bin, 10)
		word = append(word[:i], word[i+1:]...)
		i--
	  default:
		
	  }
	}
	return strings.Join(word, " ")
}
