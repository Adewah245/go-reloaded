package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(`Usage: go run . sample.txt result.txt`)
		return
	}
	text := os.Args[1]
	data, err := os.ReadFile(text)
	if err != nil {
		fmt.Println("error File Not Found!", err)
		return
	}
	result := processor(string(data))
	err = os.WriteFile(os.Args[2], []byte(result), 0644)
	fmt.Println(result)
}
