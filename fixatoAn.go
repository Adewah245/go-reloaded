package main

import "strings"

func fixatoan(s string) string {
	i := strings.ToLower(string(s[:1]))
	switch i {
	case "a", "e", "i", "o", "u", "h":
		return "an"
	}
	return "a"
}
