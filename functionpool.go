package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

type fp struct {
	fm map[string]func(string) string
}

func NewFunctionPool() *fp {
	return &fp{
		fm: make(map[string]func(string) string),
	}
}

func (f *fp) Register(name string, ptr func(string) string) *fp {
	f.fm[name] = ptr
	return f
}

func (f *fp) Run(name string, stdin io.ReadCloser, stdout, stderr io.Writer) {
	fc, exists := f.fm[name]
	if !exists {
		fmt.Fprintf(stderr, "Function %s does not exist.\n", name)
	} else {
		d, err := ioutil.ReadAll(stdin)
		if err != nil {
			fmt.Fprintf(stderr, "Failed to read from stdin: %s\n", name)
		} else {
			fmt.Fprintf(stdout, "%s", fc(string(d)))
		}
	}
}
