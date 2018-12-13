package main

import (
	"strings"
)

func upper(in string) string {
	return strings.ToUpper(in)
}

func reverse(in string) string {
	s := len(in)
	rune := make([]rune, s)
	for i, r := range in {
		rune[i] = r
	}
	rune = rune[0:s]
	for i := 0; i < s/2; i++ {
		rune[i], rune[s-1-i] = rune[s-1-i], rune[i]
	}
	return string(rune)
}
