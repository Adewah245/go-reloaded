package main

import "strings"

func cap(s string) string {
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}
