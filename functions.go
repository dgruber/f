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
