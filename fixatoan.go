package main

import (
	"strings"
	"fmt"
)

func fixatoAn(s string ) string {
	word := strings.Fields(s)
	vowel := "aeiouh"
   for i:= 0; i < len(word); i++ {
	if word[i] == "a" {
		next := word[i+1]
		word[i] = "an"
	}else{
		word[i] = "a"
	}
   }
   return strings.Join(word, " ")
}
func main(){
	fmt.Println(" god is a handsome god")
}