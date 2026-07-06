package main

import (
	"strconv"
	"strings"
)

func baseConvertion(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		if word[i] == "(hex)" && i > 0 {
			hex, err := strconv.ParseInt(word[i-1], 16, 64)
			if err != nil {
				return err.Error()
			}
			word[i-1] = strconv.FormatInt(hex, 10)
			word = append(word[:i], word[i+1:]...)
			i--
		} else if word[i] == "(bin)" && i > 0 {
			bin, err := strconv.ParseInt(word[i-1], 2, 64)
			if err != nil {
				return err.Error()
			}
			word[i-1] = strconv.FormatInt(bin, 10)
			word = append(word[:i], word[i+1:]...)
			i--
		} else if word[i] == "(octal)" && i > 0 {
			octal, err := strconv.ParseInt(word[i-1], 8, 64)
			if err != nil {
				return err.Error()
			}
			word[i-1] = strconv.FormatInt(octal, 10)
			word = append(word[:i], word[i+1:]...)
			i--
		}
	}
	return strings.Join(word, " ")
}
