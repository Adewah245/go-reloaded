package main

import "strconv"

func bin(b string) (int64, error) {
	return strconv.ParseInt(b, 2, 64)
}
