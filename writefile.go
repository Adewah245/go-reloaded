package main

import "os"

func writefile(s, w string) error {
	return os.WriteFile(s, []byte(w), 0644)
}
