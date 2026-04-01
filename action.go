package main

import "strings"

func action(s string, action string) string {
	switch action {
	case "low":
		return strings.ToLower(s)
	case "up":
		return strings.ToUpper(s)
	case "cap":
		return cap(s)
	default:
		return s
	}
}
