package main

import "strconv"

func hex(h string) (int64, error) {
	return strconv.ParseInt(h, 16, 64)
}
