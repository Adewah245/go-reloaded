package main

import (
	"os"
	"fmt"
)

func main(){
	if len(os.Args) < 3 {
		fmt.Println("usage: go run . sample.txt result.txt")
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("err can not Readfile", err)
	}

	result := processor(string(data))
	err = os.WriteFile(os.Args[2], []byte(result), 0644)
	fmt.Println(result)
}