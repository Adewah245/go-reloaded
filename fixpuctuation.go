package main

import (
	"fmt"
	"strings"
)

func fixpunc(s string ) string {
	s = strings.ReplaceAll(s, " ,", ",")
	s = strings.ReplaceAll(s, ",", ", ")
	s = strings.ReplaceAll(s, " !", "!")
	s = strings.ReplaceAll(s, " ?", "?")
	s = strings.ReplaceAll(s, " :", ":")
	s = strings.ReplaceAll(s, " .", ".")
	return s
}
func main(){
	fmt.Println(fixpunc("he ! worderful people . . ."))
	
}