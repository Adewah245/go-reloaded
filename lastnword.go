package main

import (
	"strings"
	"strconv"
)

func lastnword(s string) string {
	word := strings.Fields(s)

	for i := 0; i < len(word); i++ {
		if 1+1 < len(word) && (word[i] == "(up," || word[i] == "(low," || word[i] == "(cap,"){
			num := strings.TrimSuffix(word[i+1], ")")
			n, _ := strconv.Atoi(num)

			start := i-n
			if start < 0 {
				start = 0
			}

			for j:= start; j < i; j++{
				if word[i] == "(up," {
					if i > 0 {
					word[j] = strings.ToUpper(word[j])
					}
				} else if word[i] == "(low," {
					if i > 0 {
						word[j] = strings.ToLower(word[j])
					}
				} else if word[i] == "(cap," {
					if i > 0 && len(word[j]) > 0 {
						word[j] = strings.ToUpper(word[j][:1]) + word[j][1:]
					}
				}
			}
			word = append(word[:i], word[i+2:]...)
		i--
		}
		
	}
	return strings.Join(word, " ")
}
