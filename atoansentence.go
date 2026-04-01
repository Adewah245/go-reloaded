package main

import "strings"

func sentence(s string) string {
	s = strings.ReplaceAll(s, "A a", "An a")
	s = strings.ReplaceAll(s, "A e", "An e")
	s = strings.ReplaceAll(s, "A i", "An i")
	s = strings.ReplaceAll(s, "A o", "An o")
	s = strings.ReplaceAll(s, "A u", "An u")
	s = strings.ReplaceAll(s, "A h", "An h")
	return s
}
