package main

import "strings"

func fixspace(s []string) string {

	a := strings.Join(s, " ")
	b := strings.ReplaceAll(a, " ,", ",")
	c := strings.ReplaceAll(b, " !", "!")
	return c
}
