package main

import (
	"os"
)

func readfile(j string) (string, error) {
	f, error := os.ReadFile(j)
	if error != nil {
		return j, nil
	}
	return string(f), error
}
