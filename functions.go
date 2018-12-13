package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func upper() {
	d, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading incoming message: %s\n", err.Error())
		return
	}
	u := strings.ToUpper(string(d))
	fmt.Fprintf(os.Stdout, u)
}

func reverse() {
	d, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading incoming message: %s\n", err.Error())
		return
	}
	s := len(string(d))
	rune := make([]rune, s)
	for i, r := range string(d) {
		rune[i] = r
	}
	rune = rune[0:s]
	for i := 0; i < s/2; i++ {
		rune[i], rune[s-1-i] = rune[s-1-i], rune[i]
	}
	fmt.Fprintf(os.Stdout, string(rune))
}
