package main

import (
	"fmt"
)

func main() {
	fmt.Println(sentence("there it was. A amazine rock. A honest man."))
	fmt.Println(writefile("edge.go ", "package main\n"))
	fmt.Println(readfile("Hello.go"))
	fmt.Println(hex("2A"))
	fmt.Println(hex("7B"))
	fmt.Println(bin("11"))
	fmt.Println(bin("1101"))
	fmt.Println(cap("gOlang"))
	fmt.Println(action("I should Stop SHOUTING", "up"))
	fmt.Println(last([]string{"how", "are", "you"}, 2))
	fmt.Println(ispunc("."))
	fmt.Println(ispunc("?"))
	fmt.Println(ispunc("a"))
	fmt.Println(fixspace([]string{"hello", ",", "world", "!"}))
	fmt.Println(fixatoan("apple"), "apple")
	fmt.Println(fixatoan("book"), "book")

}
