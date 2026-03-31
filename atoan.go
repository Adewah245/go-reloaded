package main

import (
	"fmt"
	"strings"
)

func changeA(text []string) string {
	for i := 0; i < len(text)-1; i++ {
		if text[i] == "a" || text[i] == "A" {
			switch strings.ToLower(text[i+1])[0] {
			case 'a', 'e', 'i', 'o', 'u', 'h':
				text[i] = "an"
				if text[i] == "A" {
					text[i] = "An"
				}
			}
		}
	}
	return strings.Join(text, " ")
}
func main() {
	fmt.Println(changeA([]string{"there", "it", "was", "a", "amazing", "rock.", "a", "honest", "man.", "a", "book"}))

}
