package main

import "strings"

func ispunc(s string) bool {
	return strings.ContainsAny(s, ",.?:;'!")
}
